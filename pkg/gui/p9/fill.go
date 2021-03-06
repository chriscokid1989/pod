package p9

import (
	"image"

	"gioui.org/f32"
	l "gioui.org/layout"
	"gioui.org/op/paint"
)

type Filler struct {
	th  *Theme
	col string
	w   l.Widget
}

// Fill fills underneath a widget you can put over top of it
func (th *Theme) Fill(col string, w l.Widget) *Filler {
	return &Filler{th: th, col: col, w: w}
}

func (f *Filler) Embed(w l.Widget) *Filler {
	f.w = w
	return f
}

func (f *Filler) Fn(gtx l.Context) l.Dimensions {
	return f.th.Stack().Stacked(f.w).Expanded(
		func(c l.Context) l.Dimensions {
			dims := f.w(gtx)
			cs := gtx.Constraints
			d := image.Point{X: cs.Max.X, Y: cs.Max.Y}
			dr := f32.Rectangle{
				Max: f32.Point{X: float32(dims.Size.X), Y: float32(dims.Size.Y)},
			}
			paint.ColorOp{Color: f.th.Colors.Get(f.col)}.Add(gtx.Ops)
			paint.PaintOp{Rect: dr}.Add(gtx.Ops)
			gtx.Constraints.Constrain(d)
			f.w(gtx)
			gtx.Constraints.Constrain(dims.Size)
			return dims
		},
	).Fn(gtx)
}
