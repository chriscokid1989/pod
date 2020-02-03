package components

import (
	"fmt"
	"github.com/p9c/pod/cmd/gui/helpers"
	"github.com/p9c/pod/cmd/gui/models"
	"github.com/p9c/pod/cmd/gui/rcd"
	"github.com/p9c/pod/pkg/conte"
	"github.com/p9c/pod/pkg/gui/layout"
	"github.com/p9c/pod/pkg/gui/text"
	"github.com/p9c/pod/pkg/gui/unit"
)

var (
	latestTransList = &layout.List{
		Axis: layout.Vertical,
	}
)

func DuoUIlatestTxsWidget(duo *models.DuoUI, cx *conte.Xt, rc *rcd.RcVar) {

	rc.GetDuoUIlastTxs(duo, cx)

	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(duo.DuoUIcontext,
		layout.Rigid(func() {
			cs := duo.DuoUIcontext.Constraints
			helpers.DuoUIdrawRectangle(duo.DuoUIcontext, cs.Width.Max, 48, "ff3030cf", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})

			in := layout.UniformInset(unit.Dp(8))
			in.Layout(duo.DuoUIcontext, func() {

				latestx := duo.DuoUItheme.H5("Latest Transactions")
				latestx.Color = hexARGB("ffcfcfcf")
				latestx.Alignment = text.Start
				latestx.Layout(duo.DuoUIcontext)
			})
		}),
		layout.Flexed(1, func() {

			in := layout.UniformInset(unit.Dp(16))
			in.Layout(duo.DuoUIcontext, func() {
				duo.DuoUIcomponents.Status.Layout.Layout(duo.DuoUIcontext,
					// Balance status item
					layout.Rigid(func() {
						cs := duo.DuoUIcontext.Constraints
						//helpers.DuoUIdrawRectangle(duo.DuoUIcontext, cs.Width.Max, cs.Height.Max, "ff424242", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})

						//const n = 5
						//list.Layout(duo.DuoUIcontext, n, func(i int) {
						//	txt := fmt.Sprintf("List element #%d", i)
						//
						//	duo.DuoUItheme.H3(txt).Layout(duo.DuoUIcontext)
						//})
						latestTransList.Layout(duo.DuoUIcontext, len(rc.Transactions.Txs), func(i int) {
							// Invert list
							//i = len(txs.Txs) - 1 - i
							t := rc.Transactions.Txs[i]
							a := 1.0
							//const duration = 5
							helpers.DuoUIdrawRectangle(duo.DuoUIcontext, cs.Width.Max, cs.Height.Max, "ffcfcfcf", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})

							layout.Flex{
								Spacing: layout.SpaceBetween,
							}.Layout(duo.DuoUIcontext,
								layout.Rigid(func() {
									in := layout.UniformInset(unit.Dp(15))
									in.Layout(duo.DuoUIcontext, func() {

										layout.Flex{
											Axis: layout.Vertical,
										}.Layout(duo.DuoUIcontext,
											layout.Rigid(func() {

												tim := duo.DuoUItheme.Caption(t.TxID)
												tim.Color = helpers.Alpha(a, tim.Color)
												tim.Layout(duo.DuoUIcontext)
											}),
											layout.Rigid(func() {
												amount := duo.DuoUItheme.H5(fmt.Sprintf("%0.8f", t.Amount))
												amount.Color = helpers.RGB(0x003300)
												amount.Color = helpers.Alpha(a, amount.Color)
												amount.Alignment = text.End
												amount.Font.Variant = "Bold"
												amount.Font.Weight = text.Bold
												amount.Layout(duo.DuoUIcontext)
											}),
											layout.Rigid(func() {
												sat := duo.DuoUItheme.Body1(t.Category)
												sat.Color = helpers.Alpha(a, sat.Color)
												sat.Layout(duo.DuoUIcontext)
											}),
											layout.Rigid(func() {

												l := duo.DuoUItheme.Body1(helpers.FormatTime(t.Time))
												l.Color = duo.DuoUItheme.Color.Hint
												l.Color = helpers.Alpha(a, l.Color)
												l.Layout(duo.DuoUIcontext)

											}),
										)
									})

								}),
								layout.Rigid(func() {
									in := layout.UniformInset(unit.Dp(15))
									in.Layout(duo.DuoUIcontext, func() {

										sat := duo.DuoUItheme.H6(fmt.Sprintf("%0.8f", t.Amount))
										sat.Color = helpers.Alpha(a, sat.Color)
										sat.Layout(duo.DuoUIcontext)
									})

								}),
							)
						})

					}))
			})
		}),
	)
}