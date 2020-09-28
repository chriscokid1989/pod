package kopach

import (
	"context"
	"fmt"
	"github.com/stalker-loki/app/slog"
	"net"
	"os"
	"time"

	"github.com/urfave/cli"
	"go.uber.org/atomic"

	"github.com/p9c/pod/app/conte"
	"github.com/p9c/pod/cmd/kopach/client"
	"github.com/p9c/pod/cmd/kopach/control"
	"github.com/p9c/pod/cmd/kopach/control/job"
	"github.com/p9c/pod/cmd/kopach/control/pause"
	"github.com/p9c/pod/cmd/kopach/control/sol"
	chainhash "github.com/p9c/pod/pkg/chain/hash"
	"github.com/p9c/pod/pkg/comm/stdconn/worker"
	"github.com/p9c/pod/pkg/comm/transport"
	"github.com/p9c/pod/pkg/util/interrupt"
)

type HashCount struct {
	uint64
	Time time.Time
}

type Worker struct {
	active        atomic.Bool
	conn          *transport.Channel
	ctx           context.Context
	quit          chan struct{}
	cx            *conte.Xt
	sendAddresses []*net.UDPAddr
	workers       []*client.Client
	FirstSender   atomic.String
	lastSent      atomic.Int64
	Status        atomic.String
	HashTick      chan HashCount
	LastHash      *chainhash.Hash
}

func Handle(cx *conte.Xt) func(c *cli.Context) (err error) {
	return func(c *cli.Context) (err error) {
		slog.Debug("miner controller starting")
		ctx, cancel := context.WithCancel(context.Background())
		w := &Worker{
			ctx:           ctx,
			cx:            cx,
			quit:          cx.KillAll,
			sendAddresses: []*net.UDPAddr{},
		}
		w.lastSent.Store(time.Now().UnixNano())
		w.active.Store(false)
		slog.Debug("opening broadcast channel listener")
		if w.conn, err = transport.NewBroadcastChannel("kopachmain", w, *cx.Config.MinerPass,
			transport.DefaultPort, control.MaxDatagramSize, handlers, cx.KillAll); slog.Check(err) {
			cancel()
			return
		}
		var wks []*worker.Worker
		// start up the workers
		slog.Debug("starting up kopach workers")
		for i := 0; i < *cx.Config.GenThreads; i++ {
			slog.Debug("starting worker", i)
			cmd := worker.Spawn(os.Args[0], "worker",
				cx.ActiveNet.Name, *cx.Config.LogLevel)
			wks = append(wks, cmd)
			w.workers = append(w.workers, client.New(cmd.StdConn))
		}
		interrupt.AddHandler(func() {
			var err error
			w.active.Store(false)
			slog.Debug("KopachHandle interrupt")
			for i := range w.workers {
				if err = wks[i].Kill(); !slog.Check(err) {
				}
				slog.Debug("stopped worker", i)
			}
		})
		for i := range w.workers {
			slog.Debug("sending pass to worker", i)
			if err = w.workers[i].SendPass(*cx.Config.MinerPass); slog.Check(err) {
			}
		}
		w.active.Store(true)
		// controller watcher thread
		go func() {
			slog.Debug("starting controller watcher")
			ticker := time.NewTicker(time.Second)
		out:
			for {
				select {
				case <-ticker.C:
					// if the last message sent was 3 seconds ago the server is
					// almost certainly disconnected or crashed so clear
					// FirstSender
					since := time.Now().Sub(time.Unix(0, w.lastSent.Load()))
					wasSending := since > time.Second*3 && w.FirstSender.Load() != ""
					if wasSending {
						slog.Debug("previous current controller has stopped"+
							" broadcasting", since, w.FirstSender.Load())
						// when this string is clear other broadcasts will be listened to
						w.FirstSender.Store("")
						// pause the workers
						for i := range w.workers {
							slog.Debug("sending pause to worker", i)
							if err = w.workers[i].Pause(); slog.Check(err) {
							}
						}
					}
				case <-cx.KillAll:
					break out
				}
			}
		}()
		slog.Debug("listening on", control.UDP4MulticastAddress)
		<-cx.KillAll
		slog.Info("kopach shutting down")
		return
	}
}

// these are the handlers for specific message types.
var handlers = transport.Handlers{
	string(job.Magic): func(ctx interface{}, src net.Addr, dst string,
		b []byte) (err error) {
		w := ctx.(*Worker)
		if !w.active.Load() {
			slog.Debug("not active")
			return
		}
		j := job.LoadContainer(b)
		ips := j.GetIPs()
		cP := j.GetControllerListenerPort()
		addr := net.JoinHostPort(ips[0].String(), fmt.Sprint(cP))
		firstSender := w.FirstSender.Load()
		otherSent := firstSender != addr && firstSender != ""
		if otherSent {
			slog.Debug("ignoring other controller job")
			// ignore other controllers while one is active and received first
			return
		}
		if firstSender == "" {
			slog.Warn("new sender", addr)
		}
		w.FirstSender.Store(addr)
		w.lastSent.Store(time.Now().UnixNano())
		for i := range w.workers {
			if err = w.workers[i].NewJob(&j); slog.Check(err) {
			}
		}
		return
	},
	string(pause.Magic): func(ctx interface{}, src net.Addr, dst string, b []byte) (err error) {
		w := ctx.(*Worker)
		p := pause.LoadPauseContainer(b)
		fs := w.FirstSender.Load()
		ni := p.GetIPs()[0].String()
		np := p.GetControllerListenerPort()
		ns := net.JoinHostPort(ni, fmt.Sprint(np))
		if fs == ns {
			for i := range w.workers {
				slog.Debug("sending pause to worker", i, fs, ns)
				if err = w.workers[i].Pause(); slog.Check(err) {
				}
			}
		}
		return
	},
	string(sol.SolutionMagic): func(ctx interface{}, src net.Addr, dst string,
		b []byte) (err error) {
		w := ctx.(*Worker)
		// port := strings.Split(w.FirstSender.Load(), ":")[1]
		// j := sol.LoadSolContainer(b)
		// senderPort := j.GetSenderPort()
		// if fmt.Sprint(senderPort) == port {
		// 	Warn("we found a solution")
		// }
		w.FirstSender.Store("")
		return
	},
}
