package sampler

import "github.com/sporsh/gosundog/geometry"

type PathTraceSampler struct {
	Geometry geometry.Intersectable
	Epsilon  float64
}

func (pt PathTraceSampler) Sample(r geometry.Ray) (s Sample) {
	if _, ok := pt.Geometry.Intersect(r, pt.Epsilon); ok {
		s[0], s[1], s[2] = 1, 1, 1
	}
	return
}
