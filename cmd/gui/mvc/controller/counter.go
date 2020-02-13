// SPDX-License-Identifier: Unlicense OR MIT

package controller

import (
	"time"

	"github.com/p9c/pod/pkg/gui/gesture"
	"github.com/p9c/pod/pkg/gui/layout"
	"github.com/p9c/pod/pkg/gui/op"
)

type Counter struct {
	click gesture.Click
	// clicks tracks the number of unreported clicks.
	clicks int
	// prevClicks tracks the number of unreported clicks
	// that belong to the previous frame.
	prevClicks int
	history    []Click
}

//// Click represents a historic click.
//type Click struct {
//	Position f32.Point
//	Time     time.Time
//}

func (b *Counter) Clicked(gtx *layout.Context) bool {
	b.processEvents(gtx)
	if b.clicks > 0 {
		b.clicks--
		if b.prevClicks > 0 {
			b.prevClicks--
		}
		if b.clicks > 0 {
			// Ensure timely delivery of remaining clicks.
			op.InvalidateOp{}.Add(gtx.Ops)
		}
		return true
	}
	return false
}

func (b *Counter) History() []Click {
	return b.history
}

func (b *Counter) Layout(gtx *layout.Context) {
	// Flush clicks from before the previous frame.
	b.clicks -= b.prevClicks
	b.prevClicks = 0
	b.processEvents(gtx)
	b.click.Add(gtx.Ops)
	for len(b.history) > 0 {
		c := b.history[0]
		if gtx.Now().Sub(c.Time) < 1*time.Second {
			break
		}
		copy(b.history, b.history[1:])
		b.history = b.history[:len(b.history)-1]
	}
}

func (b *Counter) processEvents(gtx *layout.Context) {
	for _, e := range b.click.Events(gtx) {
		switch e.Type {
		case gesture.TypeClick:
			b.clicks++
		case gesture.TypePress:
			b.history = append(b.history, Click{
				Position: e.Position,
				Time:     gtx.Now(),
			})
		}
	}
}