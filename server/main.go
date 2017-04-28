package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"

	sundog "github.com/sporsh/gosundog"
	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/sampler"
	"github.com/sporsh/gosundog/v3"
)

type PathTraceImage struct {
	sampler.PathTraceSampler
	rect image.Rectangle
}

func NewPathTraceImage(g geometry.Intersectable, c sundog.Camera, width, height int) *PathTraceImage {
	return &PathTraceImage{
		sampler.PathTraceSampler{g, 0.0001},
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
		c := sundog.Camera{
			v3.V{0, 0, -400},
			geometry.Basis{v3.Y, v3.X, v3.Z},
			1.0,
		}
		img := NewPathTraceImage(g, c, 640, 480)
		log.Println("###", img)
		if err := png.Encode(w, img); err != nil {
			fmt.Fprintln(w, err)
		}
	})
	log.Println("Starting server")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
