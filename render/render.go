package main

import (
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	sundog "github.com/sporsh/gosundog"
	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

func main() {
	rand.Seed(0)

	// lightMaterial := sundog.NewLambertianMaterial(v3.V{0.5, 0.5, 2}, v3.V{0.8, 0.8, 0.8})
	lightMaterial := sundog.NewLambertianMaterial(v3.V{2, 2, 2}, v3.V{0, 0, 0})
	whiteMaterial := sundog.NewLambertianMaterial(v3.ZERO, v3.V{0.8, 0.8, 0.8})
	mirrorMaterial := sundog.NewSpecularMaterial(v3.ZERO, v3.V{0.8, 0.8, 0.8})
	glassMaterial := sundog.NewRefractiveMaterial(v3.ZERO, v3.V{0.8, 0.8, 0.8})
	redMaterial := sundog.NewLambertianMaterial(v3.ZERO, v3.V{0.75, 0.25, 0.25})
	greenMaterial := sundog.NewLambertianMaterial(v3.ZERO, v3.V{0.25, 0.75, 0.25})

	g := &geometry.Group{
		// sundog.Renderable{
		// 	Intersectable: geometry.NewSphere(v3.V{-0.55, -0.5, 0}, 0.25),
		// 	Material:      redMaterial,
		// },

		sundog.Renderable{
			Intersectable: geometry.NewSphere(v3.V{-0.5, -0.5, 0}, .3),
			Material:      glassMaterial,
		},
		sundog.Renderable{
			Intersectable: geometry.NewSphere(v3.V{0, 0.15, 0.3}, .3),
			Material:      mirrorMaterial,
		},
		sundog.Renderable{
			Intersectable: geometry.NewSphere(v3.V{0.5, -0.5, -1}, .3),
			Material:      lightMaterial,
		},

		// sundog.Renderable{
		// 	Intersectable: geometry.NewTorus88(0.4, 0.4),
		// 	// Intersectable: geometry.NewDSphere(0.6),
		// 	Material:      whiteMaterial,
		// },

		// sundog.Renderable{
		// 	Intersectable: geometry.NewSphere(v3.V{0.55, -0.5, 0}, 0.25),
		// 	Material:      greenMaterial,
		// },

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
		// sundog.Renderable{
		// 	Intersectable: geometry.NewSphere(v3.V{0, 1, 0}, 0.25),
		// 	Material:      lightMaterial,
		// },
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
		FocalLength: 1.5,
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	img := sundog.NewPathTraceImage(g, c, 400, 400, 50, 6)

	t0 := time.Now()
	png.Encode(f, img)
	log.Println(time.Since(t0))

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
