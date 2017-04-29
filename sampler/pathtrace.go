package sampler

import (
	sundog "github.com/sporsh/gosundog"
	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type PathTraceSampler struct {
	Geometry geometry.Intersectable
	Epsilon  float64
}

func (pt PathTraceSampler) Sample(r geometry.Ray) (s Sample) {
	if i, ok := pt.Geometry.Intersect(r, pt.Epsilon); ok {
		if obj, ok := i.Geometry.(sundog.Renderable); ok {
			out := v3.Negate(r.Direction)

			s.Radiance = v3.Add(
				s.Radiance,
				obj.Material.Emittance(out, i.Normal),
			)

			// weight := obj.Material.ReflectancePDF(out, i.Normal)
			// prob := math.Max(math.Max(weight[0], weight[1]), weight[2])

			// in := obj.Material.Reflect(out, i.Normal)

		}
	}
	return
}
