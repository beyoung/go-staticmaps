package sm

import (
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	"image/color"
	"testing"
)

func TestStaticMap(t *testing.T) {
	ctx := NewContext()
	ctx.SetSize(400, 300)
	ctx.AddObject(
		NewMarker(
			s2.LatLngFromDegrees(52.514536, 13.350151),
			color.RGBA{R: 0xff, A: 0xff},
			16.0,
		),
	)
	ctx.AddObject(
		NewMarker(
			s2.LatLngFromDegrees(22.514536, 13.350151),
			color.RGBA{R: 0xff, A: 0xff},
			16.0,
		))

	ctx.AddObject(
		NewMarker(
			s2.LatLngFromDegrees(22.514536, 113.350151),
			color.RGBA{R: 0xff, A: 0xff},
			16.0,
		))
	ctx.AddObject(
		NewMarker(
			s2.LatLngFromDegrees(-22.514536, 113.350151),
			color.RGBA{R: 0xff, A: 0xff},
			16.0,
		))

	img, err := ctx.Render()
	if err != nil {
		panic(err)
	}

	if err := gg.SavePNG("my-map.png", img); err != nil {
		panic(err)
	}
}
