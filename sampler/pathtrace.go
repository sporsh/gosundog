package sampler

import (
	"math"
	"math/rand"

	sundog "github.com/sporsh/gosundog"
	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type PathTraceSampler struct {
	Geometry geometry.Intersectable
	Epsilon  float64
}

func (pt PathTraceSampler) Sample(r geometry.Ray) Sample {
	s := Sample{
		Radiance: v3.ZERO,
		Weight:   v3.V{1, 1, 1},
	}
	maxBounces := 4
	for bounces := 0; bounces < maxBounces; bounces++ {
		if i, ok := pt.Geometry.Intersect(r, pt.Epsilon); ok {
			if obj, ok := i.Geometry.(sundog.Renderable); ok {
				out := v3.Negate(r.Direction)

				s.Radiance = v3.Add(
					s.Radiance,
					v3.Hadamard(
						s.Weight,
						obj.Material.Emittance(out, i.Basis),
					),
				)

				// Russian roulette

				// The probability for light to be reflected in this direction
				prob := obj.Material.ReflectancePDF(out, i.Basis)
				pMax := math.Max(math.Max(prob[0], prob[1]), prob[2])

				if bounces > 2 {
					if rand.Float64() <= pMax {
						prob = v3.Scale(prob, 1/pMax)
					} else {
						break
					}
				}

				s.Weight = v3.Hadamard(
					prob,
					s.Weight,
				)

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
