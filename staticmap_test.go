package sm

import (
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	"image"
	"image/color"
	"os"
	"testing"
)

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func TestStaticMap(t *testing.T) {
	ctx := NewContext()
	ctx.SetSize(1280, 560)
	start, _ := getImageFromFilePath("./start.png")

	logo, _ := getImageFromFilePath("./logo.png")

	ctx.SetLogo(Logo{
		Logo: logo,
		X:    1.1,
		Y:    1.3,
	})

	ctx.AddObject(
		NewImageMarker(s2.LatLngFromDegrees(52.524536, 13.350151),
			start, 16, 16,
		),
	)

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
