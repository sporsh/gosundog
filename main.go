package main

import (
	"fmt"
	"image/color"
	"image/png"
	"net/http"

	"github.com/sporsh/gosundog/geometry"
)

type PathTraceImage struct {
}

func (img PathTraceImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img PathTraceImage) At(x, y int) color.Color {
	r := geometry.Ray{}
	return img.Sampler.Sample(r, img.Sampler.Epsilon)
}

func NewPathTraceImage() *PathTraceImage {
	return &PathTraceImage{}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		img := NewPathTraceImage()

		if err := png.Encode(w, img); err != nil {
			fmt.Fprintln(w, err)
		}
	})
}
