package geometry_test

import (
	"math"
	"testing"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

func TestSphereIntersect(t *testing.T) {
	r := geometry.Ray{
		Origin:    v3.V{0, 0, 0},
		Direction: v3.V{0, 0, 1},
		TMin:      0,
		TMax:      math.Inf(1),
	}
	s := geometry.NewSphere(
		v3.V{0, 0, 1},
		1,
	)

	if i, ok := s.Intersect(&r); !ok {
		t.Log(i)
		t.Fail()
	}
}

// func BenchmarkAdd(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_ = v3.V{1, 2, 3}
// 	}
// }
