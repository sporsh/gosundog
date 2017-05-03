package geometry

import (
	"math"

	"github.com/sporsh/gosundog/v3"
)

type Basis struct {
	Tangent, Bitangent, Normal v3.V
}

func (b Basis) ToLocal(v v3.V) v3.V {
	return v3.V{
		v3.Dot(&v, &b.Tangent),
		v3.Dot(&v, &b.Bitangent),
		v3.Dot(&v, &b.Normal),
	}
}

func (b Basis) ToWorld(v v3.V) v3.V {
	return v3.V{
		b.Tangent.X*v.X + b.Normal.X*v.Y + b.Bitangent.X*v.Z,
		b.Tangent.Y*v.X + b.Normal.Y*v.Y + b.Bitangent.Y*v.Z,
		b.Tangent.Z*v.X + b.Normal.Z*v.Y + b.Bitangent.Z*v.Z,
	}
}

func ArbritraryBasisForNormal(normal v3.V) Basis {
	tangent := OrthogonalUnitVector(normal)
	return Basis{
		Tangent:   tangent,
		Bitangent: v3.Cross(tangent, normal),
		Normal:    normal,
	}
}

func OrthogonalUnitVector(v v3.V) v3.V {
	if v.X == 0 {
		return v3.V{1, 0, 0}
	}
	f := math.Sqrt(v.X*v.X + v.Z*v.Z)
	return v3.V{v.Z * f, 0, -v.X * f}
}
