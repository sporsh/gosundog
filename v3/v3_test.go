package v3_test

import (
	"testing"

	"github.com/sporsh/gosundog/v3"
)

func TestAdd(t *testing.T) {
	a := v3.V{0.1, 0.2, 0.3}
	b := v3.V{-1, -2, -3}

	t.Run("a+a", func(t *testing.T) {
		actual := v3.Add(a, a)
		expected := v3.V{0.2, 0.4, 0.6}
		if actual != expected {
			t.Logf("%#v != %#v", actual, expected)
			t.Fail()
		}
	})

	t.Run("a+b", func(t *testing.T) {
		actual := v3.Add(a, b)
		expected := v3.V{-0.9, -1.8, -2.7}
		if actual != expected {
			t.Logf("%#v != %#v", actual, expected)
			t.Fail()
		}
	})

	// t.Run("Chain", func(t *testing.T) {
	// 	actual := v3.Add(a, v3.Add(a, b))
	// 	expected := v3.V{-0.8, -1.6, -2.4}
	// 	if actual != expected {
	// 		t.Logf("%#v != %#v", actual, expected)
	// 		t.Fail()
	// 	}
	// })

}

type VS struct {
	X, Y, Z float64
}

func (a VS) Add(b VS) VS {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
	return a
}

type VA [3]float64

func (a VA) Add(b VA) VA {
	a[0] += b[0]
	a[1] += b[1]
	a[2] += b[2]
	return a
}

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = VS{1, 2, 3}.Add(VS{4, 5, 6})
	}
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = VA{1, 2, 3}.Add(VA{4, 5, 6})
	}
}
