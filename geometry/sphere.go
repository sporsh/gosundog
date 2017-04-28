package geometry

import (
	"math"

	"github.com/sporsh/gosundog/v3"
)

type Sphere struct {
	Center  v3.V
	Radius  float64
	radius2 float64
}

func NewSphere(c v3.V, r float64) Sphere {
	return Sphere{c, r, r * r}
}

func (s Sphere) Intersect(r Ray, epsilon float64) (i Intersection, ok bool) {
	m := v3.Sub(r.Origin, s.Center)
	c := v3.Len2(m) - s.radius2

	b := v3.Dot(m, r.Direction)
	if b > 0 {
		return i, false
	}

	discr := b*b - c
	if discr < 0 {
		return i, false
	}

	// inside := false
	sqrtDiscr := math.Sqrt(discr)
	i.T = -b - sqrtDiscr
	if i.T < epsilon {
		// inside = true
		i.T = -b + sqrtDiscr
		if i.T < epsilon {
			return i, false
		}
	}

	i.Point = v3.Add(
		r.Origin,
		v3.Scale(r.Direction, i.T),
	)

	i.Normal = v3.Normalize(v3.Sub(
		i.Point,
		s.Center,
	))

	return i, true
}
