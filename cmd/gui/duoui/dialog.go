package duoui

import (
	"github.com/p9c/pod/cmd/gui/helpers"
	"github.com/p9c/pod/cmd/gui/rcd"
	"github.com/p9c/pod/pkg/conte"
	"github.com/p9c/pod/pkg/gui/layout"
	"github.com/p9c/pod/pkg/gui/text"
	"github.com/p9c/pod/pkg/gui/unit"
	"github.com/p9c/pod/pkg/gui/widget"
	"github.com/p9c/pod/pkg/gui/widget/parallel"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image/color"
)

var (
	buttonDialogCancel = new(widget.Button)
	buttonDialogOK     = new(widget.Button)
	buttonDialogClose  = new(widget.Button)
)

// Main wallet screen
func (duo *DuoUI) DuoUIdialog(cx *conte.Xt, rc *rcd.RcVar) {
	// START View <<<
	//cs := duo.m.DuoUIcontext.Constraints
	//	helpers.DuoUIdrawRectangle(duo.m.DuoUIcontext, cs.Width.Max, cs.Height.Max, "ee303030", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
	//layout.Flexed(1, func() {
	//
	//	layout.Align(layout.Center).Layout(duo.m.DuoUIcontext, func() {
	//		layout.Inset{Top: unit.Dp(24), Bottom: unit.Dp(8), Left: unit.Dp(0), Right: unit.Dp(4)}.Layout(duo.m.DuoUIcontext, func() {
	//			cur := duo.m.DuoUItheme.H4("dddddddddddddddddddddddddddd")
	//			cur.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0xcf}
	//			cur.Alignment = text.Start
	//			cur.Layout(duo.m.DuoUIcontext)
	//		})
	//	})
	//
	//})
	iconCancel, _ := parallel.NewDuoUIicon(icons.NavigationCancel)
	iconOK, _ := parallel.NewDuoUIicon(icons.NavigationCheck)
	iconClose, _ := parallel.NewDuoUIicon(icons.NavigationClose)

	cs := duo.m.DuoUIcontext.Constraints
	helpers.DuoUIdrawRectangle(duo.m.DuoUIcontext, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ee000000"), [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
	layout.Align(layout.Center).Layout(duo.m.DuoUIcontext, func() {
		//cs := duo.m.DuoUIcontext.Constraints
		helpers.DuoUIdrawRectangle(duo.m.DuoUIcontext, 408, 150, duo.m.DuoUItheme.Color.Primary, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})

		layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		}.Layout(duo.m.DuoUIcontext,
			layout.Rigid(func() {
				layout.Flex{
					Axis:      layout.Horizontal,
					Alignment: layout.Middle,
				}.Layout(duo.m.DuoUIcontext,
					layout.Rigid(func() {
						layout.Align(layout.Center).Layout(duo.m.DuoUIcontext, func() {
							layout.Inset{Top: unit.Dp(24), Bottom: unit.Dp(8), Left: unit.Dp(0), Right: unit.Dp(4)}.Layout(duo.m.DuoUIcontext, func() {
								cur := duo.m.DuoUItheme.H4("DIALOG BOX!")
								cur.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0xcf}
								cur.Alignment = text.Start
								cur.Layout(duo.m.DuoUIcontext)
							})
						})
					}),
				)
			}),
			layout.Rigid(func() {
				layout.Flex{
					Axis:      layout.Horizontal,
					Alignment: layout.Middle,
				}.Layout(duo.m.DuoUIcontext,
					layout.Rigid(dialogButon("CANCEL", "ffcfcfcf", "ffcf3030", "ffcfcfcf", duo, rc, buttonDialogCancel, iconCancel)),
					layout.Rigid(dialogButon("OK", "ffcfcfcf", "ff308030", "ffcfcfcf", duo, rc, buttonDialogOK, iconOK)),
					layout.Rigid(dialogButon("CLOSE", "ffcfcfcf", "ffcf8030", "ffcfcfcf", duo, rc, buttonDialogClose, iconClose)),

				)
			}),
		)
	})
}

func dialogButon(text, txtColor, bgColor, iconColor string, duo *DuoUI, rc *rcd.RcVar, button *widget.Button, icon *parallel.DuoUIicon) func() {
	var b parallel.DuoUIbutton
	return func() {
		layout.Inset{Top: unit.Dp(8), Bottom: unit.Dp(8), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(duo.m.DuoUIcontext, func() {
			b = duo.m.DuoUItheme.DuoUIbutton(text, txtColor, bgColor, iconColor, 24, 120, 60, 0, 0, icon)
			for button.Clicked(duo.m.DuoUIcontext) {
				rc.ShowDialog = false
			}
			b.Layout(duo.m.DuoUIcontext, button)
		})
	}
}
