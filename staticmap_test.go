package sm

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"
	"os"
	"testing"

	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
)

var (
	//go:embed start.png
	startIcon []byte
	start     image.Image

	//go:embed end.png
	endIcon []byte
	end     image.Image
)

func init() {
	start, _, _ = image.Decode(bytes.NewReader(startIcon))
	end, _, _ = image.Decode(bytes.NewReader(endIcon))
}

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

func TestLineMap(t *testing.T) {
	points := []s2.LatLng{
		s2.LatLngFromDegrees(31.291278, 121.517232),
		s2.LatLngFromDegrees(31.291286, 121.517264),
		s2.LatLngFromDegrees(31.29128, 121.51723),
		s2.LatLngFromDegrees(31.291307, 121.5172),
		s2.LatLngFromDegrees(31.29145, 121.51693),
	}
	startPoint := points[0]
	endpoint := points[len(points)-1]

	ctx := NewContext()

	// ctx.SetTileProvider(NewTileProviderXingzheInternal())

	ctx.SetSize(800, 600)

	ctx.AddObject(NewPath(points, color.Black, 2))

	offsetX, offsetY := 16.0, 16.0

	ctx.AddObject(
		NewImageMarker(
			s2.LatLngFromDegrees(startPoint.Lat.Degrees(), startPoint.Lng.Degrees()),
			start,
			offsetX, offsetY),
	)
	ctx.AddObject(
		NewImageMarker(
			s2.LatLngFromDegrees(endpoint.Lat.Degrees(), endpoint.Lng.Degrees()),
			end,
			offsetX, offsetY),
	)
	staticmap, err := ctx.Render()
	if err != nil {
		log.Fatalln(err)
	}
	if err := gg.SavePNG("workout.png", staticmap); err != nil {
		panic(err)
	}

}
