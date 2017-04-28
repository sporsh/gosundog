package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/sampler"
	"github.com/sporsh/gosundog/v3"
)

type PathTraceImage struct {
	sampler.PathTraceSampler
	rect image.Rectangle
}

func NewPathTraceImage(g geometry.Intersectable, c Camera, width, height int) *PathTraceImage {
	return &PathTraceImage{
		sampler.PathTraceSampler{},
		image.Rect(0, 0, width, height),
	}
}

func (img PathTraceImage) At(x, y int) color.Color {
	r := geometry.Ray{}
	return img.Sample(r)
}

func (img PathTraceImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img PathTraceImage) Bounds() image.Rectangle {
	return img.rect
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		g := geometry.Group{geometry.NewSphere(v3.V{0, 0, 0}, 100)}
		c := NewCamera()
		img := NewPathTraceImage(g, c, 640, 480)

		if err := png.Encode(w, img); err != nil {
			fmt.Fprintln(w, err)
		}
	})
}
