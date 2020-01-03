// SPDX-License-Identifier: Unlicense OR MIT

package theme

import (
	"image/color"

	"github.com/p9c/pod/pkg/gio/layout"
	"github.com/p9c/pod/pkg/gio/op/paint"
	"github.com/p9c/pod/pkg/gio/text"
	"github.com/p9c/pod/pkg/gio/unit"
	"github.com/p9c/pod/cmd/gui/widget"
)

type DuoUIlabel struct {
	// Face defines the text style.
	Font text.Font
	// Color is the text color.
	Color color.RGBA
	// Alignment specify the text alignment.
	Alignment text.Alignment
	// MaxLines limits the number of lines. Zero means no limit.
	MaxLines int
	Text     string

	shaper *text.Shaper
}

func (t *DuoUItheme) H1(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(96.0/16.0), txt)
}

func (t *DuoUItheme) H2(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(60.0/16.0), txt)
}

func (t *DuoUItheme) H3(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(48.0/16.0), txt)
}

func (t *DuoUItheme) H4(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(34.0/16.0), txt)
}

func (t *DuoUItheme) H5(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(24.0/16.0), txt)
}

func (t *DuoUItheme) H6(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(20.0/16.0), txt)
}

func (t *DuoUItheme) Body1(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize, txt)
}

func (t *DuoUItheme) Body2(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(14.0/16.0), txt)
}

func (t *DuoUItheme) Caption(txt string) DuoUIlabel {
	return t.DuoUIlabel(t.TextSize.Scale(12.0/16.0), txt)
}

func (t *DuoUItheme) DuoUIlabel(size unit.Value, txt string) DuoUIlabel {
	return DuoUIlabel{
		Text:  txt,
		Color: t.Color.Text,
		Font: text.Font{
			Size: size,
		},
		shaper: t.Shaper,
	}
}

func (l DuoUIlabel) Layout(gtx *layout.Context) {
	paint.ColorOp{Color: l.Color}.Add(gtx.Ops)
	tl := widget.Label{Alignment: l.Alignment, MaxLines: l.MaxLines}
	tl.Layout(gtx, l.shaper, l.Font, l.Text)
}