package sundog

import "github.com/sporsh/gosundog/v3"

type Material interface {
	BRDF(in, out, normal v3.V) (pd v3.V)
	Emittance(out, normal v3.V) v3.V
	Reflect(out, normal v3.V) (in v3.V)
	ReflectancePDF(out, normal v3.V) (pd v3.V)
}

type LambertianMaterial struct {
	RadiantEmittance v3.V
}

func (m LambertianMaterial) BRDF(in, out, normal v3.V) v3.V {
	return v3.ZERO
}

// Emittance returns the radiant emittance (or exitance) in a given outgoing direction
func (m LambertianMaterial) Emittance(out, normal v3.V) v3.V {
	return m.RadiantEmittance
}

func (m LambertianMaterial) Reflect(out, normal v3.V) v3.V {
	return v3.X
}

func (m LambertianMaterial) ReflectancePDF(out, normal v3.V) v3.V {
	return v3.ZERO
}
