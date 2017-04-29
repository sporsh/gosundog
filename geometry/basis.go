package geometry

import "github.com/sporsh/gosundog/v3"

type Basis struct {
	Tangent, Bitangent, Normal v3.V
}

func (b Basis) ToLocal(v v3.V) v3.V {
	return v3.V{
		v3.Dot(v, b.Tangent),
		v3.Dot(v, b.Bitangent),
		v3.Dot(v, b.Normal),
	}
}

func (b Basis) ToWorld(v v3.V) v3.V {
	return v3.V{
		b.Tangent[0]*v[0] + b.Bitangent[0]*v[1] + b.Normal[0]*v[2],
		b.Tangent[1]*v[0] + b.Bitangent[1]*v[1] + b.Normal[1]*v[2],
		b.Tangent[2]*v[0] + b.Bitangent[2]*v[1] + b.Normal[2]*v[2],
	}
}
