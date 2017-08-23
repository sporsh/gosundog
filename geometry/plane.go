package geometry

import "github.com/sporsh/gosundog/v3"

type Plane struct {
	Normal v3.V
	D      float64
	basis  Basis
}

func NewPlane(n v3.V, d float64) *Plane {
	basis := ArbritraryBasisForNormal(n)
	return &Plane{
		Normal: n,
		D:      d,
		basis:  basis,
	}
}

func (p *Plane) Intersect(r *Ray) (i Intersection, ok bool) {
	i.T = (p.D - v3.Dot(p.Normal, r.Origin)) / v3.Dot(p.Normal, r.Direction)
	if i.T > r.TMin && i.T < r.TMax {
		i.Point = v3.Add(r.Origin, v3.Scale(r.Direction, i.T))
		i.Basis = p.basis
		return i, true
	}
	return i, false
}
