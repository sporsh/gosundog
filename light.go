package sundog

import (
	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type Light interface {
}

type PointLight struct {
	Origin          v3.V
	RadiantExitance v3.V
}

func (l *PointLight) Ray(p v3.V) *geometry.Ray {
	ld := v3.Sub(l.Origin, p)

	return &geometry.Ray{
		Direction: v3.Normalize(ld),
		Origin:    p,
		TMax:      v3.Len(ld),
		TMin:      0,
	}
}
