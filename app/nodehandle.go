package app

import (
	"github.com/stalker-loki/app/slog"
	"github.com/stalker-loki/pod/app/config"
	"github.com/urfave/cli"

	"github.com/stalker-loki/pod/app/conte"
	"github.com/stalker-loki/pod/cmd/node"
)

func nodeHandle(cx *conte.Xt) func(c *cli.Context) error {
	return func(c *cli.Context) (err error) {
		slog.Trace("running node handler")
		config.Configure(cx, c.Command.Name, true)
		cx.NodeReady = make(chan struct{})
		cx.Node.Store(false)
		// serviceOptions defines the configuration options for the daemon as a service on Windows.
		type serviceOptions struct {
			ServiceCommand string `short:"s" long:"service" description:"Service command {install, remove, start, stop}"`
		}
		// runServiceCommand is only set to a real function on Windows.  It is used to parse and execute service
		// commands specified via the -s flag.
		var runServiceCommand func(string) error
		// Service options which are only added on Windows.
		serviceOpts := serviceOptions{}
		// Perform service command and exit if specified.  Invalid service commands show an appropriate error.
		// Only runs on Windows since the runServiceCommand function will be nil when not on Windows.
		if serviceOpts.ServiceCommand != "" && runServiceCommand != nil {
			err := runServiceCommand(serviceOpts.ServiceCommand)
			if err != nil {
				slog.Error(err)
				return err
			}
			return nil
		}
		shutdownChan := make(chan struct{})
		go func() {
			err := node.Main(cx, shutdownChan)
			if err != nil {
				slog.Error("error starting node ", err)
			}
		}()
		slog.Debug("sending back node rpc server handler")
		cx.RPCServer = <-cx.NodeChan
		close(cx.NodeReady)
		cx.Node.Store(true)
		cx.WaitGroup.Wait()
		return nil
	}
}
