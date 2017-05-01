package sundog

import (
	"math"
	"math/rand"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type Material interface {
	BRDF(in, out v3.V, basis geometry.Basis) (pd v3.V)
	Emittance(out v3.V, basis geometry.Basis) v3.V
	Reflect(out v3.V, basis geometry.Basis) (in v3.V)
	ReflectancePDF(out v3.V, basis geometry.Basis) (pd v3.V)
}

type LambertianMaterial struct {
	RadiantEmittance v3.V
	Reflectivity     v3.V
}

func (m LambertianMaterial) BRDF(in, out v3.V, basis geometry.Basis) v3.V {
	return *m.Reflectivity.Scale(math.Pi)
}

// Emittance returns the radiant emittance (or exitance) in a given outgoing direction
func (m LambertianMaterial) Emittance(out v3.V, basis geometry.Basis) v3.V {
	return m.RadiantEmittance
}

// Reflect returns a incoming direction that would reflect to the given outgoing direction
func (m LambertianMaterial) Reflect(out v3.V, basis geometry.Basis) v3.V {
	u1, u2 := rand.Float64(), rand.Float64()

	sinTheta := math.Sqrt(u1)
	cosTheta := math.Sqrt(1 - u1)

	phi := 2 * math.Pi * u2

	return v3.Normalize(basis.ToWorld(v3.V{
		sinTheta * math.Cos(phi),
		cosTheta,
		sinTheta * math.Sin(phi),
	}))
}

// ReflectancePDF calculates the probability for a ray to be reflected
// in the given outgoing direction
func (m LambertianMaterial) ReflectancePDF(out v3.V, basis geometry.Basis) v3.V {
	return m.Reflectivity
}
