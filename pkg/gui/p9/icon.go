package p9

import (
	"image"
	"image/color"
	"image/draw"

	"gioui.org/f32"
	l "gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"golang.org/x/exp/shiny/iconvg"
)

type Icon struct {
	th    *Theme
	color color.RGBA
	src   []byte
	size  unit.Value
	// Cached values.
	sz       int
	op       paint.ImageOp
	imgSize  int
	imgColor color.RGBA
}

// Icon returns a new Icon from iconVG data.
func (th *Theme) Icon() *Icon {
	return &Icon{th: th, size: th.TextSize, color: rgb(0xff000000)}
}

// Color sets the color of the icon image. It must be called before creating the image
func (i *Icon) Color(color string) *Icon {
	i.color = i.th.Colors.Get(color)
	return i
}

func (i *Icon) RGBA(rgba color.RGBA) *Icon {
	i.color = rgba
	return i
}

// Src sets the icon source to draw from
func (i *Icon) Src(data []byte) *Icon {
	_, err := iconvg.DecodeMetadata(data)
	if Check(err) {
		return nil
	}
	i.src = data
	return i
}

// Scale changes the size relative to the base font size
func (i *Icon) Scale(scale float32) *Icon {
	i.size = i.th.TextSize.Scale(scale * 1)
	return i
}

func (i *Icon) Size(size unit.Value) *Icon {
	i.size = size
	return i
}

// Fn renders the icon
func (i *Icon) Fn(gtx l.Context) l.Dimensions {
	ico := i.image(gtx.Px(i.size))
	ico.Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: toPointF(ico.Size()),
		},
	}.Add(gtx.Ops)
	return l.Dimensions{Size: ico.Size()}
}

func (i *Icon) image(sz int) paint.ImageOp {
	if sz == i.imgSize && i.color == i.imgColor {
		return i.op
	}
	m, _ := iconvg.DecodeMetadata(i.src)
	dx, dy := m.ViewBox.AspectRatio()
	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz,
		Y: int(float32(sz) * dy / dx)}})
	var ico iconvg.Rasterizer
	ico.SetDstImage(img, img.Bounds(), draw.Src)
	m.Palette[0] = i.color
	iconvg.Decode(&ico, i.src, &iconvg.DecodeOptions{
		Palette: &m.Palette,
	})
	i.op = paint.NewImageOp(img)
	i.imgSize = sz
	i.imgColor = i.color
	return i.op
}
