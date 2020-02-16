package duoui

import (
	"fmt"
	"github.com/p9c/pod/cmd/gui/mvc/theme"
	"github.com/p9c/pod/pkg/gui/clipboard"
	"github.com/p9c/pod/pkg/gui/layout"
	"github.com/p9c/pod/pkg/gui/unit"
)

var (
	addressBookList = &layout.List{
		Axis:        layout.Vertical,
		//ScrollToEnd: true,
	}
)

func (ui *DuoUI) DuoUIaddressBook() func() {
	return func() {
		layout.Flex{}.Layout(ui.ly.Context,
			layout.Flexed(1, func() {
				layout.UniformInset(unit.Dp(0)).Layout(ui.ly.Context, func() {
					layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceAround,
					}.Layout(ui.ly.Context,
						layout.Flexed(1, func() {
							addressBookList.Layout(ui.ly.Context, len(ui.rc.AddressBook.Addresses), func(i int) {
								t := ui.rc.AddressBook.Addresses[i]
								layout.Flex{Axis: layout.Vertical}.Layout(ui.ly.Context,
									layout.Rigid(func() {
										layout.Flex{
											Alignment: layout.End,
										}.Layout(ui.ly.Context,
											layout.Flexed(0.1, func() {
												sat := ui.ly.Theme.Body1(fmt.Sprint(t.Index))
												sat.Font.Typeface = ui.ly.Theme.Font.Primary
												sat.Color = theme.HexARGB(ui.ly.Theme.Color.Dark)
												sat.Layout(ui.ly.Context)
											}),
											layout.Rigid(func() {

												var copyButton theme.DuoUIbutton
												copyButton = ui.ly.Theme.DuoUIbutton(ui.ly.Theme.Font.Mono, t.Address, ui.ly.Theme.Color.Light, ui.ly.Theme.Color.Primary, "", "", 0, 300, 24, 0, 0)

												for t.Copy.Clicked(ui.ly.Context) {

													clipboard.Set(t.Address)
												}
												copyButton.Layout(ui.ly.Context, t.Copy)

											}),
											layout.Flexed(0.2, func() {
												sat := ui.ly.Theme.Body1(t.Account)
												sat.Font.Typeface = ui.ly.Theme.Font.Primary
												sat.Color = theme.HexARGB(ui.ly.Theme.Color.Dark)
												sat.Layout(ui.ly.Context)
											}),
											layout.Flexed(0.4, func() {
												sat := ui.ly.Theme.Body1(t.Label)
												sat.Font.Typeface = ui.ly.Theme.Font.Primary
												sat.Color = theme.HexARGB(ui.ly.Theme.Color.Dark)
												sat.Layout(ui.ly.Context)
											}),
											layout.Flexed(0.3, func() {
												sat := ui.ly.Theme.Body1(fmt.Sprint(t.Amount))
												sat.Font.Typeface = ui.ly.Theme.Font.Primary
												sat.Color = theme.HexARGB(ui.ly.Theme.Color.Dark)
												sat.Layout(ui.ly.Context)
											}),

										)
									}),
									layout.Rigid(func() {
										cs := ui.ly.Context.Constraints
										theme.DuoUIdrawRectangle(ui.ly.Context, cs.Width.Max, 1, ui.ly.Theme.Color.Hint, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
									}))
							})
						}))
				})
			}),
		)
	}
}
