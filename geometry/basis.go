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
		b.Tangent[0]*v[0] + b.Normal[0]*v[1] + b.Bitangent[0]*v[2],
		b.Tangent[1]*v[0] + b.Normal[1]*v[1] + b.Bitangent[1]*v[2],
		b.Tangent[2]*v[0] + b.Normal[2]*v[1] + b.Bitangent[2]*v[2],
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
	if v[0] == 0 {
		return v3.V{1, 0, 0}
	}
	f := math.Sqrt(v[0]*v[0] + v[2]*v[2])
	return v3.V{v[2] * f, 0, -v[0] * f}
}
