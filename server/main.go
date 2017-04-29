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
	rect image.Rectangle
	sampler.PathTraceSampler
	camera sundog.Camera
}

func NewPathTraceImage(g geometry.Intersectable, c sundog.Camera, width, height int) *PathTraceImage {
	return &PathTraceImage{
		rect: image.Rect(0, 0, width, height),
		PathTraceSampler: sampler.PathTraceSampler{
			Geometry: g,
			Epsilon:  0.0001,
		},
		camera: c,
	}
}

func (img PathTraceImage) At(x, y int) color.Color {
	r := img.camera.RayThrough(x, y)
	if x == 25 && y == 25 {
		log.Printf("%#v/n", r)
	}
	return img.Sample(r)
}

func (img PathTraceImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img PathTraceImage) Bounds() image.Rectangle {
	return img.rect
}

func main() {
	width, height := 100, 100

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		g := geometry.Group{geometry.NewSphere(v3.V{50, 50, 50}, 50)}
		c := sundog.Camera{
			Origin: v3.V{50, 50, -50},
			Basis: geometry.Basis{
				X: v3.Y,
				Y: v3.X,
				Z: v3.Z,
			},
			Aperture: 1,
		}
		img := NewPathTraceImage(g, c, width, height)
		log.Println("Starting...")
		if err := png.Encode(w, img); err != nil {
			fmt.Fprintln(w, err)
		}
		log.Println("Done!")
	})
	log.Println("Starting server")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
