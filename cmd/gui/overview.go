package gui

import (
	"fmt"
	"strings"

	l "gioui.org/layout"
	icons2 "golang.org/x/exp/shiny/materialdesign/icons"

	"github.com/p9c/pod/pkg/gui/p9"
)

func (wg *WalletGUI) OverviewPage() l.Widget {
	return func(gtx l.Context) l.Dimensions {
		balanceColumn := wg.th.Column(p9.Rows{
			{Label: "Available:", W: wg.balanceWidget(wg.State.balance)},
			{Label: "Unconfirmed:", W: wg.balanceWidget(wg.State.balanceUnconfirmed)},
			{Label: "Total:", W: wg.balanceWidget(wg.State.balance + wg.State.balanceUnconfirmed)},
		}, "bariol bold", 1).List
		return wg.th.Responsive(*wg.App.Size, p9.Widgets{
			{
				Widget: wg.th.VFlex().
					Rigid(
						func(gtx l.Context) l.Dimensions {
							_, bc := balanceColumn(gtx)
							return wg.th.Inset(0.25,
								wg.th.Fill("DocBg",
									wg.th.SliceToWidget(
										append([]l.Widget{
											func(gtx l.Context) l.Dimensions {
												_, bc = balanceColumn(gtx)
												gtx.Constraints.Max.X = *wg.Size
												return wg.th.Fill("PanelBg",
													wg.th.Flex().
														Rigid(
															// wg.th.Inset(0.5,
															wg.th.H6("Balances").
																// Font("bariol bold").
																Color("PanelText").
																Fn,
															// ).Fn,
														).Fn,
												).Fn(gtx)
											},
										},
											bc...), l.Vertical),
								).Fn,
							).Fn(gtx)
						},
					).
					Flexed(1,
						wg.th.Inset(0.25,
							wg.th.VFlex().Rigid(
								wg.Fill("PanelBg",
									wg.th.Flex().
										Rigid(
											// wg.Inset(0.5,
											wg.H6("recent transactions").Color("PanelText").Fn,
										).Fn,
									// ).Fn,
								).
									Fn,
							).Rigid(
								wg.th.Fill("DocBg",
									wg.th.Inset(0.25,
										wg.th.Flex().
											// Rigid(
											// 	wg.th.Inset(0.25,
											// 		p9.EmptySpace(0, 0)).Fn,
											// ).
											Flexed(1,
												wg.RecentTransactions(),
											).Fn,
									).Fn,
								).Fn,
							).Fn,
						).Fn,
					).
					Fn,
			},
			{
				Size: 1280,
				Widget: wg.th.Flex().
					Rigid(
						func(gtx l.Context) l.Dimensions {
							_, bc := balanceColumn(gtx)
							return wg.th.Inset(0.25,
								wg.th.Fill("DocBg",
									wg.th.SliceToWidget(
										append([]l.Widget{
											func(gtx l.Context) l.Dimensions {
												// render the widgets onto a second context to get their dimensions
												gtx1 := p9.CopyContextDimensions(gtx, gtx.Constraints.Max, l.Vertical)
												dim := p9.GetDimension(gtx1, wg.th.SliceToWidget(bc, l.Vertical))
												gtx.Constraints.Max.X = dim.Size.X
												gtx.Constraints.Min.X = dim.Size.X
												return wg.th.Fill("PanelBg",
													wg.th.Flex().
														Flexed(1,
															// wg.th.Inset(0.5,
															wg.th.H6("Balances").
																// Font("bariol bold").
																Color("PanelText").
																Fn,
															// ).Fn,
														).Fn,
												).Fn(gtx)
											},
										},
											bc...), l.Vertical),
								).Fn,
							).Fn(gtx)
						},
					).
					Flexed(1,
						wg.th.Inset(0.25,
							wg.th.VFlex().Rigid(
								wg.Fill("PanelBg",
									wg.th.Flex().
										Rigid(
											// wg.Inset(0.5,
											wg.H6("recent transactions").Color("DocText").Fn,
											// ).Fn,
										).Fn,
								).
									Fn,
							).Rigid(
								wg.th.Fill("DocBg",
									wg.th.Inset(0.25,
										wg.th.Flex().
											// Rigid(
											// 	wg.th.Inset(0.25,
											// 		p9.EmptySpace(0, 0)).Fn,
											// ).
											Flexed(1,
												wg.RecentTransactions(),
											).Fn,
									).Fn,
								).Fn,
							).Fn,
						).Fn,
					).
					Fn,
			},
		}).Fn(gtx)
	}
}

// RecentTransactions generates a display showing recent transactions
//
// fields to use: Address, Amount, BlockIndex, BlockTime, Category, Confirmations, Generated
func (wg *WalletGUI) RecentTransactions() l.Widget {
	var out []l.Widget
	first := true
	// out = append(out)
	for x := range wg.State.lastTxs {
		i := x
		txs := wg.State.lastTxs[i]
		// spacer
		if !first {
			out = append(out,
				wg.th.Fill("DocBg",
					wg.th.Inset(0.25, p9.EmptyMaxWidth()).Fn,
				).Fn,
			)
		} else {
			first = false
		}
		out = append(out,
			wg.th.Fill("DocBg",
				wg.th.Body1(fmt.Sprintf("%-6.8f DUO", txs.Amount)).Color("PanelText").Fn,
			).Fn,
		)

		out = append(out,
			wg.th.Fill("DocBg",
				wg.th.Caption(txs.Address).
					Font("go regular").
					Color("PanelText").
					TextScale(0.66).Fn,
			).Fn,
		)

		out = append(out,
			wg.th.Fill("DocBg",
				wg.th.Caption(txs.TxID).
					Font("go regular").
					Color("PanelText").
					TextScale(0.5).Fn,
			).Fn,
		)
		out = append(out,
			func(gtx l.Context) l.Dimensions {
				return wg.th.Fill("DocBg",
					wg.th.Flex().AlignMiddle(). // SpaceBetween().
						Rigid(
							wg.th.Flex().AlignMiddle().
								Rigid(
									wg.Icon().Color("DocText").Scale(1).Src(&icons2.DeviceWidgets).Fn,
								).
								// Rigid(
								// 	wg.th.Caption(fmt.Sprint(*txs.BlockIndex)).Fn,
								// 	// wg.buttonIconText(txs.clickBlock,
								// 	// 	fmt.Sprint(*txs.BlockIndex),
								// 	// 	&icons2.DeviceWidgets,
								// 	// 	wg.blockPage(*txs.BlockIndex)),
								// ).
								Rigid(
									wg.th.Caption(fmt.Sprintf("%d ", *txs.BlockIndex)).Fn,
								).
								Fn,
						).
						Rigid(
							wg.th.Flex().AlignMiddle().
								Rigid(
									wg.Icon().Color("DocText").Scale(1).Src(&icons2.ActionCheckCircle).Fn,
								).
								Rigid(
									wg.th.Caption(fmt.Sprintf("%d ", txs.Confirmations)).Fn,
								).
								Fn,
						).
						Rigid(
							wg.th.Flex().AlignMiddle().
								Rigid(
									func(gtx l.Context) l.Dimensions {
										switch txs.Category {
										case "generate":
											return wg.Icon().Color("DocText").Scale(1).Src(&icons2.ActionStars).Fn(gtx)
										case "immature":
											return wg.Icon().Color("DocText").Scale(1).Src(&icons2.ImageTimeLapse).Fn(gtx)
										case "receive":
											return wg.Icon().Color("DocText").Scale(1).Src(&icons2.ActionPlayForWork).Fn(gtx)
										case "unknown":
											return wg.Icon().Color("DocText").Scale(1).Src(&icons2.AVNewReleases).Fn(gtx)
										}
										return l.Dimensions{}
									},
								).
								Rigid(
									wg.th.Caption(txs.Category+" ").Fn,
								).
								Fn,
						).
						Rigid(
							wg.th.Flex().AlignMiddle().
								Rigid(
									wg.Icon().Color("DocText").Scale(1).Src(&icons2.DeviceAccessTime).Fn,
								).
								Rigid(
									wg.th.Caption(
										wg.State.lastTimeStrings[i],
										// txs.time,
									).Color("DocText").Fn,
								).
								Fn,
						).
						// TODO: this thing hasn't got data going in yet, before we can display anything we need data
						//  also the index `i` is not from wg.State.txs it is from wg.State.lastTxs
						//  - even if these two data sets overlap if you want them to relate to each other you need
						//  to define their integration. Simple way would be for eg: as you intend, to merge them into
						//  one and only update (add) the extra data on page display. I think that it's so trivial for
						//  10 instances of the listtransactions result struct just keep them separate so the logic is
						//  cleaner. In other words, add a second fetcher in ticker.go for the history/tx page, and
						//  handle the damn empty list, nil panics are Satan.
						//  my advice is use the second field and keep them separated
						// Rigid(
						// 	wg.Inset(0.1, wg.buttonText(wg.State.txs[i].clickTx, "details", wg.txPage(i))).Fn,
						// ).
						Fn,
				).
					Fn(gtx)
			})
	}
	le := func(gtx l.Context, index int) l.Dimensions {
		return out[index](gtx)
	}
	return func(gtx l.Context) l.Dimensions {
		return wg.lists["recent"].
			Vertical().
			Length(len(out)).
			ListElement(le).
			Fn(gtx)
	}
}

func leftPadTo(length, limit int, txt string) string {
	if len(txt) > limit {
		return txt[limit-len(txt):]
	}
	pad := length - len(txt)
	return strings.Repeat(" ", pad) + txt
}

func (wg *WalletGUI) balanceWidget(balance float64) l.Widget {
	bal := leftPadTo(15, 15, fmt.Sprintf("%6.8f", balance))
	return wg.th.Flex().AlignEnd().
		Rigid(wg.th.Body1(" ").Fn).
		Rigid(
			wg.th.Caption(bal).
				Font("go regular").
				Fn,
		).
		Fn
}

//
// func (wg *WalletGUI) panel(title string, fill bool, content l.Widget) l.Widget {
// 	return func(gtx l.Context) l.Dimensions {
// 		w := wg.Inset(0.25,
// 			wg.Fill("DocBg",
// 				wg.th.VFlex().
// 					Rigid(
// 						wg.Fill("DocText",
// 							wg.th.Flex().
// 								Rigid(
// 									wg.Inset(0.5,
// 										wg.H6(title).Color("DocBg").Fn,
// 									).Fn,
// 								).Fn,
// 						).Fn,
// 					).
// 					Rigid(
// 						wg.Fill("DocBg",
// 							wg.Inset(0.25,
// 								content,
// 							).Fn,
// 						).Fn,
// 					).Fn,
// 			).Fn,
// 		).Fn
// 		if !fill {
// 			// render the widgets onto a second context to get their dimensions
// 			gtx1 := p9.CopyContextDimensions(gtx, gtx.Constraints.Max, l.Vertical)
// 			// generate the dimensions for all the list elements
// 			child := op.Record(gtx1.Ops)
// 			d := w(gtx1)
// 			_ = child.Stop()
// 			gtx.Constraints.Max.X = d.Size.X
// 			gtx.Constraints.Max.Y = d.Size.Y
// 			gtx.Constraints.Min = gtx.Constraints.Max
// 			w = wg.Inset(0.25,
// 				wg.th.VFlex().
// 					Rigid(
// 						wg.Fill("DocText",
// 							wg.th.Flex().
// 								Flexed(1,
// 									wg.Inset(0.5,
// 										wg.H6(title).Color("DocBg").Fn,
// 									).Fn,
// 								).Fn,
// 						).Fn,
// 					).
// 					Rigid(
// 						wg.Fill("DocBg",
// 							wg.Inset(0.25,
// 								content,
// 							).Fn,
// 						).Fn,
// 					).Fn,
// 			).Fn
// 		}
// 		return w(gtx)
// 	}
// }
