package sundog

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type PathTraceImage struct {
	PathTraceSampler
	rect       image.Rectangle
	camera     Camera
	numSamples int
}

func NewPathTraceImage(g geometry.Intersectable, c Camera, width, height, numSamples, maxBounces int) *PathTraceImage {
	return &PathTraceImage{
		rect: image.Rect(0, 0, width, height),
		PathTraceSampler: PathTraceSampler{
			Geometry:   g,
			Epsilon:    0.001,
			maxBounces: maxBounces,
		},
		camera:     c,
		numSamples: numSamples,
	}
}

func (img *PathTraceImage) At(x, y int) color.Color {
	radiance := v3.V{X: 0, Y: 0, Z: 0}
	numSamples := img.numSamples
	for sample := 0; sample < numSamples; sample++ {
		u := 2*(float64(x)+rand.Float64())/float64(img.rect.Dx()-1) - 1
		v := 1 - 2*(float64(y)+rand.Float64())/float64(img.rect.Dy()-1)
		r := img.camera.RayThrough(u, v, img.Epsilon)
		sampleRadiance := img.Sample(&r).Radiance
		radiance = v3.Add(radiance, sampleRadiance)
	}
	return Sample{
		Radiance: v3.Scale(radiance, 1.0/float64(numSamples)),
	}
}

func (img *PathTraceImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *PathTraceImage) Bounds() image.Rectangle {
	return img.rect
}

func (img *PathTraceImage) Opaque() bool {
	return true
}

type PathTraceSampler struct {
	Geometry   geometry.Intersectable
	Epsilon    float64
	maxBounces int
}

func (pt *PathTraceSampler) Sample(r *geometry.Ray) Sample {
	s := Sample{
		Radiance: v3.ZERO,
		Weight:   v3.V{X: 1, Y: 1, Z: 1},
	}
	maxBounces := pt.maxBounces
	for bounces := 0; bounces < maxBounces; bounces++ {
		if i, ok := pt.Geometry.Intersect(r); ok {
			if obj, ok := i.Geometry.(Renderable); ok {
				out := v3.Negate(r.Direction)

				//  Add the emittance of the intersected object's material
				emittance := obj.Material.Emittance(out, i.Basis)
				s.Radiance = v3.Add(s.Radiance, v3.Hadamard(emittance, s.Weight))

				//  Add light paths from point lights
				ld := v3.Sub(v3.V{0, 0.9, 0}, i.Point)
				lr := &geometry.Ray{
					Direction: v3.Normalize(ld),
					Origin:    i.Point,
					TMax:      v3.Len(ld) - pt.Epsilon,
					TMin:      pt.Epsilon,
				}
				if _, ok := pt.Geometry.Intersect(lr); !ok {
					in := lr.Direction
					s.Radiance = v3.Add(
						s.Radiance,
						v3.Hadamard(
							s.Weight,
							v3.Hadamard(
								obj.Material.BRDF(in, out, i.Basis),
								v3.Scale(v3.V{1, 1, 1}, 1/v3.Len2(ld)),
							),
						),
					)
				}

				// The probability for light to be reflected in this direction
				prob := obj.Material.ReflectancePDF(out, i.Basis)
				pMax := math.Max(math.Max(prob.X, prob.Y), prob.Z)

				// Russian roulette (after a couple of bounces)
				if rand.Float64() > pMax {
					break
				}
				prob = v3.Scale(prob, 1/pMax)

				s.Weight = v3.Hadamard(s.Weight, prob)

				in := obj.Material.Reflect(out, i.Basis)
				r.Direction = in
				r.Origin = i.Point
			}
		} else {
			break
		}
	}
	return s
}
