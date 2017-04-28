package sundog

import (
	"math"
	"math/rand"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type Camera struct {
	Origin   v3.V
	Basis    geometry.Basis
	Aperture float64
}

func (c Camera) LookAt(p v3.V) {}

func (c Camera) RandomOriginWithinAperture() v3.V {
	u1, u2 := rand.Float64(), rand.Float64()
	r := u2 * c.Aperture
	localPoint := v3.V{
		math.Cos(u1*2*math.Pi) * r,
		math.Sin(u1*2*math.Pi) * r,
		0,
	}
	return v3.Add(
		c.Basis.ToLocal(localPoint),
		c.Origin,
	)
}
