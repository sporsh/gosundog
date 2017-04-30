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
			Epsilon:  0.001,
		},
		camera: c,
	}
}

func (img PathTraceImage) At(x, y int) color.Color {
	u := 2*float64(x)/float64(img.rect.Dx()-1) - 1
	v := 1 - 2*float64(y)/float64(img.rect.Dy()-1)
	radiance := v3.V{0, 0, 0}
	numSamples := 100
	for sample := 0; sample < numSamples; sample++ {
		r := img.camera.RayThrough(u, v)
		radiance = v3.Add(
			radiance,
			img.Sample(r).Radiance,
		)
	}
	return sampler.Sample{
		Radiance: v3.Scale(radiance, 1.0/float64(numSamples)),
	}
}

func (img PathTraceImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img PathTraceImage) Bounds() image.Rectangle {
	return img.rect
}

func main() {
	width, height := 300, 300

	lightMaterial := sundog.LambertianMaterial{
		RadiantEmittance: v3.V{20, 20, 20},
	}

	whiteMaterial := sundog.LambertianMaterial{
		// RadiantEmittance: v3.V{.1, .1, .1},
		Reflectivity: v3.V{0.8, 0.8, 0.8},
	}
	redMaterial := sundog.LambertianMaterial{
		Reflectivity: v3.V{0.75, 0.25, 0.25},
	}
	greenMaterial := sundog.LambertianMaterial{
		// RadiantEmittance: v3.V{.1, .1, .1},
		Reflectivity: v3.V{0.25, 0.75, 0.25},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		g := geometry.Group{
			sundog.Renderable{
				Intersectable: geometry.NewSphere(v3.V{0, -0.25, 0}, .25),
				Material:      whiteMaterial,
			},

			// Cornell box
			sundog.Renderable{
				Intersectable: geometry.NewPlane(v3.V{1, 0, 0}, -1),
				Material:      redMaterial,
			},
			sundog.Renderable{
				Intersectable: geometry.NewPlane(v3.V{-1, 0, 0}, -1),
				Material:      greenMaterial,
			},
			sundog.Renderable{
				Intersectable: geometry.NewPlane(v3.V{0, 0, -1}, -1),
				Material:      whiteMaterial,
			},
			sundog.Renderable{
				Intersectable: geometry.NewPlane(v3.V{0, 1, 0}, -1),
				Material:      whiteMaterial,
			},
			sundog.Renderable{
				Intersectable: geometry.NewPlane(v3.V{0, -1, 0}, -1),
				Material:      whiteMaterial,
			},
			// Light
			sundog.Renderable{
				Intersectable: geometry.NewSphere(v3.V{0, 1, 0}, 0.2),
				Material:      lightMaterial,
			},
		}
		c := sundog.Camera{
			Origin: v3.V{0, 0, -1.5},
			Basis: geometry.Basis{
				Tangent:   v3.X,
				Bitangent: v3.Y,
				Normal:    v3.Z,
			},
			Aperture:    0.05,
			FieldOfView: 1,
			FocalLength: 2,
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
