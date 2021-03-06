package p9

import (
	"image"
	"image/color"

	"gioui.org/f32"
	l "gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/p9c/pod/pkg/gui/f32color"
)

type Slider struct {
	th       *Theme
	min, max float32
	color    color.RGBA
	float    *Float
}

// Slider is for selecting a value in a range.
func (th *Theme) Slider() *Slider {
	return &Slider{
		th:    th,
		color: th.Colors.Get("Primary"),
	}
}

// Min sets the value at the left hand side
func (s *Slider) Min(min float32) *Slider {
	s.min = min
	return s
}

// Max sets the value at the right hand side
func (s *Slider) Max(max float32) *Slider {
	s.max = max
	return s
}

// Color sets the color to draw the slider in
func (s *Slider) Color(color string) *Slider {
	s.color = s.th.Colors.Get(color)
	return s
}

// Float sets the initial value
func (s *Slider) Float(f *Float) *Slider {
	s.float = f
	return s
}

// Fn renders the slider
func (s *Slider) Fn(c l.Context) l.Dimensions {
	thumbRadiusInt := c.Px(unit.Sp(6))
	trackWidth := float32(c.Px(unit.Sp(2)))
	thumbRadius := float32(thumbRadiusInt)
	halfWidthInt := 2 * thumbRadiusInt
	halfWidth := float32(halfWidthInt)

	size := c.Constraints.Max
	// Keep a minimum length so that the track is always visible.
	minLength := halfWidthInt + 3*thumbRadiusInt + halfWidthInt
	if size.X < minLength {
		size.X = minLength
	}
	size.Y = 2 * halfWidthInt

	st := op.Push(c.Ops)
	op.Offset(f32.Pt(halfWidth, 0)).Add(c.Ops)
	c.Constraints.Min = image.Pt(size.X-2*halfWidthInt, size.Y)
	s.float.Fn(c, halfWidthInt, s.min, s.max)
	thumbPos := halfWidth + s.float.Pos()
	st.Pop()

	color := s.color
	if c.Queue == nil {
		color = f32color.MulAlpha(color, 150)
	}

	// Draw track before thumb.
	st = op.Push(c.Ops)
	track := f32.Rectangle{
		Min: f32.Point{
			X: halfWidth,
			Y: halfWidth - trackWidth/2,
		},
		Max: f32.Point{
			X: thumbPos,
			Y: halfWidth + trackWidth/2,
		},
	}
	clip.RRect{Rect: track}.Add(c.Ops)
	paint.ColorOp{Color: color}.Add(c.Ops)
	paint.PaintOp{Rect: track}.Add(c.Ops)
	st.Pop()

	// Draw track after thumb.
	st = op.Push(c.Ops)
	track.Min.X = thumbPos
	track.Max.X = float32(size.X) - halfWidth
	clip.RRect{Rect: track}.Add(c.Ops)
	paint.ColorOp{Color: f32color.MulAlpha(color, 96)}.Add(c.Ops)
	paint.PaintOp{Rect: track}.Add(c.Ops)
	st.Pop()

	// Draw thumb.
	st = op.Push(c.Ops)
	thumb := f32.Rectangle{
		Min: f32.Point{
			X: thumbPos - thumbRadius,
			Y: halfWidth - thumbRadius,
		},
		Max: f32.Point{
			X: thumbPos + thumbRadius,
			Y: halfWidth + thumbRadius,
		},
	}
	rr := thumbRadius
	clip.RRect{
		Rect: thumb,
		NE:   rr, NW: rr, SE: rr, SW: rr,
	}.Add(c.Ops)
	paint.ColorOp{Color: color}.Add(c.Ops)
	paint.PaintOp{Rect: thumb}.Add(c.Ops)
	st.Pop()

	return l.Dimensions{Size: size}
}
