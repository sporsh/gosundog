package geometry

import "github.com/sporsh/gosundog/v3"

type Intersectable interface {
	Intersect(r *Ray, epsilon float64) (i Intersection, ok bool)
}

type Intersection struct {
	T        float64
	Point    v3.V
	Basis    Basis
	Geometry Intersectable
}
