package geometry

import "github.com/sporsh/gosundog/v3"

type Basis struct {
	X, Y, Z v3.V
}

func (b Basis) ToLocal(v v3.V) v3.V {
	return v3.V{
		v3.Dot(v, b.X),
		v3.Dot(v, b.Y),
		v3.Dot(v, b.Z),
	}
}

func (b Basis) ToWorld(v v3.V) v3.V {
	return v3.V{
		b.X[0]*v[0] + b.Y[0]*v[1] + b.Z[0]*v[2],
		b.X[1]*v[0] + b.Y[1]*v[1] + b.Z[1]*v[2],
		b.X[2]*v[0] + b.Y[2]*v[1] + b.Z[2]*v[2],
	}
}
