package p9

import (
	"fmt"

	l "gioui.org/layout"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type IncDec struct {
	th                *Theme
	nDigits           int
	min, max          int
	Current           int
	changeHook        func(n int)
	inc, dec          *Clickable
	color, background string
	inactive          string
}

// IncDec is a simple increment/decrement for a number setting
func (th *Theme) IncDec(nDigits, min, max, current int, changeHook func(n int)) (out *IncDec) {
	out = &IncDec{
		th:         th,
		nDigits:    nDigits,
		min:        min,
		max:        max,
		Current:    current,
		changeHook: changeHook,
		inc:        th.Clickable(),
		dec:        th.Clickable(),
		// color:      color,
		// background: background,
		// inactive:   inactive,
	}
	return
}

func (in *IncDec) SetColor(color string) *IncDec {
	in.color = color
	return in
}

func (in *IncDec) SetBackground(color string) *IncDec {
	in.background = color
	return in
}
func (in *IncDec) SetInactive(color string) *IncDec {
	in.inactive = color
	return in
}

func (in *IncDec) Fn(gtx l.Context) l.Dimensions {
	out := in.th.Flex().AlignMiddle()
	incColor, decColor := in.color, in.color
	if in.Current == in.min {
		decColor = in.inactive
	}
	if in.Current == in.max {
		incColor = in.inactive
	}
	if in.Current == in.min {
		out.Rigid(
			in.th.Inset(0.25,
				in.th.Icon().Color(decColor).Scale(Scales["H5"]).Src(&icons.ContentRemove).Fn,
			).Fn,
		)
	} else {
		out.Rigid(in.th.Inset(0.25,
			in.th.ButtonLayout(in.inc.SetClick(func() {
				in.Current--
				if in.Current < in.min {
					in.Current = in.min
				} else {
					in.changeHook(in.Current)
				}
			})).Background("Transparent").Embed(
				in.th.Icon().Color(decColor).Scale(Scales["H5"]).Src(&icons.ContentRemove).Fn,
			).Fn,
		).Fn,
		)
	}
	cur := fmt.Sprintf("%"+fmt.Sprint(in.nDigits)+"d", in.Current)
	out.Rigid(in.th.Body1(cur).Color(in.color).Font("go regular").Fn)
	if in.Current == in.max {
		out.Rigid(
			in.th.Inset(0.25,
				in.th.Icon().Color(incColor).Scale(Scales["H5"]).Src(&icons.ContentAdd).Fn,
			).Fn,
		)
	} else {
		out.Rigid(
			in.th.Inset(0.25,
				in.th.ButtonLayout(in.dec.SetClick(func() {
					in.Current++
					if in.Current > in.max {
						in.Current = in.max
					} else {
						in.changeHook(in.Current)
					}
				})).Background("Transparent").Embed(
					in.th.Icon().Color(incColor).Scale(Scales["H5"]).Src(&icons.ContentAdd).Fn,
				).Fn,
			).Fn,
		)
	}
	return out.Fn(gtx)
}