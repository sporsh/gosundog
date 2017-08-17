package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"

	sundog "github.com/sporsh/gosundog"
	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

func main() {
	width, height := 400, 400

	lightMaterial := sundog.LambertianMaterial{
		RadiantEmittance: v3.V{20, 20, 20},
	}

	whiteMaterial := sundog.LambertianMaterial{
		Reflectivity: v3.V{0.8, 0.8, 0.8},
	}
	redMaterial := sundog.LambertianMaterial{
		Reflectivity: v3.V{0.75, 0.25, 0.25},
	}
	greenMaterial := sundog.LambertianMaterial{
		Reflectivity: v3.V{0.25, 0.75, 0.25},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		g := geometry.Group{

			sundog.Renderable{
				Intersectable: geometry.NewSphere(v3.V{-0.55, -0.5, 0}, 0.25),
				Material:      redMaterial,
			},
			sundog.Renderable{
				Intersectable: geometry.NewSphere(v3.V{0, -0.5, 0.25}, .25),
				Material:      whiteMaterial,
			},
			sundog.Renderable{
				Intersectable: geometry.NewSphere(v3.V{0.55, -0.5, 0}, 0.25),
				Material:      greenMaterial,
			},
			// Cornell box
			// sundog.Renderable{
			// 	Intersectable: geometry.NewPlane(v3.V{1, 0, 0}, -1),
			// 	Material:      redMaterial,
			// },
			// sundog.Renderable{
			// 	Intersectable: geometry.NewPlane(v3.V{-1, 0, 0}, -1),
			// 	Material:      greenMaterial,
			// },
			// sundog.Renderable{
			// 	Intersectable: geometry.NewPlane(v3.V{0, 0, -1}, -1),
			// 	Material:      whiteMaterial,
			// },
			sundog.Renderable{
				Intersectable: geometry.NewPlane(v3.V{0, 1, 0}, -1),
				Material:      whiteMaterial,
			},
			// sundog.Renderable{
			// 	Intersectable: geometry.NewPlane(v3.V{0, -1, 0}, -1),
			// 	Material:      whiteMaterial,
			// },
			// Light

			sundog.Renderable{
				Intersectable: geometry.NewSphere(v3.V{0, 1, 0}, 0.25),
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
			Aperture: 0,
			// Aperture:    0.05,
			FieldOfView: 1,
			FocalLength: 2,
		}
		img := sundog.NewPathTraceImage(&g, c, width, height)
		log.Println("Starting...")
		if err := png.Encode(w, img); err != nil {
			fmt.Fprintln(w, err)
		}
		log.Println("Done!")
	})
	log.Println("Starting server")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
