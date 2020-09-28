// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
)

type DuoUIThemeNav struct {
	Title string
	// Color is the text color.
	TxColor      color.RGBA
	Font         text.Font
	BgColor      color.RGBA
	CornerRadius unit.Value
	// Icon          *DuoUIIcon
	IcoBackground color.RGBA
	IcoColor      color.RGBA
	IcoPadding    unit.Value
	IcoSize       unit.Value
	Size          unit.Value
	Padding       unit.Value
	// NavButtons    map[string]*DuoUIbutton
	// theme         DuoUITheme
}

// func (t *DuoUITheme) DuoUIThemeNav(Txt string, items *map[string]*DuoUIbutton) DuoUIThemeNav {
//	//for it, item := range items {
//	//	items[it] = t.DuoUIbutton(item.Text, item.Icon)
//	//}
//
//	return DuoUIThemeNav{
//		Title: Txt,
//		Font: text.Font{
//			Size: t.TextSize.Scale(14.0 / 16.0),
//		},
//		BgColor:    t.Color.Primary,
//		TxColor:    t.Color.InvText,
//		Size:       unit.Dp(56),
//		Padding:    unit.Dp(16),
//		NavButtons: *items,
//		theme:      *t,
//	}
// }

func (n DuoUIThemeNav) Layout(gtx *layout.Context) {
	// navList := &layout.List{
	//	Axis: layout.Vertical,
	// }
	//
	// navButtons := make(map[int]layout.Widget)
	//
	// for a, b := range n.NavButtons {
	//	navButtons[b.Order] = func() {
	//		n.theme.H3(a).Layout(gtx)
	//	}
	// }
	// //for a, _ := range n.NavButtons {
	// //	navButtons = append(navButtons, func() {
	// //		n.theme.H3(a).Layout(gtx)
	// //	})
	// //}
	// //	navButtons := func() {
	// //		n.theme.H3("button").Layout(gtx)
	// //	}
	//
	// navList.Layout(gtx, 2, func(i int) {
	//	layout.UniformInset(unit.Dp(16)).Layout(gtx, navButtons[i])
	// })
}
