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
	emittance    v3.V
	reflectivity v3.V
	brdf         v3.V
}

func NewLambertianMaterial(emittance, reflectivity v3.V) *LambertianMaterial {
	return &LambertianMaterial{
		emittance:    emittance,
		reflectivity: reflectivity,
		brdf:         v3.Scale(reflectivity, 1/math.Pi),
	}
}

func (m *LambertianMaterial) BRDF(in, out v3.V, basis geometry.Basis) v3.V {
	return v3.Scale(m.brdf, math.Max(0, v3.Dot(basis.Normal, in)))
}

// Emittance returns the radiant emittance (or exitance) in a given outgoing direction
func (m *LambertianMaterial) Emittance(out v3.V, basis geometry.Basis) v3.V {
	return m.emittance
}

// Reflect returns a incoming direction that would reflect to the given outgoing direction
func (m *LambertianMaterial) Reflect(out v3.V, basis geometry.Basis) v3.V {
	u1, u2 := rand.Float64(), rand.Float64()

	sinTheta := math.Sqrt(u1)
	cosTheta := math.Sqrt(1 - u1)

	phi := 2 * math.Pi * u2

	return v3.Normalize(basis.ToWorld(v3.V{
		X: sinTheta * math.Cos(phi),
		Y: cosTheta,
		Z: sinTheta * math.Sin(phi),
	}))
}

// ReflectancePDF calculates the probability for a ray to be reflected
// in the given outgoing direction
func (m LambertianMaterial) ReflectancePDF(out v3.V, basis geometry.Basis) v3.V {
	return m.reflectivity
}

type SpecularMaterial struct {
	emittance    v3.V
	reflectivity v3.V
}

func NewSpecularMaterial(emittance, reflectivity v3.V) *SpecularMaterial {
	return &SpecularMaterial{
		emittance:    emittance,
		reflectivity: reflectivity,
	}
}

func (m *SpecularMaterial) BRDF(in, out v3.V, basis geometry.Basis) v3.V {
	return v3.ZERO
}

func (m *SpecularMaterial) Emittance(out v3.V, basis geometry.Basis) v3.V {
	return m.emittance
}

func (m *SpecularMaterial) Reflect(out v3.V, basis geometry.Basis) v3.V {
	return v3.Sub(
		v3.Scale(basis.Normal, 2*math.Max(0, v3.Dot(basis.Normal, out))),
		out,
	)
}

func (m SpecularMaterial) ReflectancePDF(out v3.V, basis geometry.Basis) v3.V {
	return m.reflectivity
}
