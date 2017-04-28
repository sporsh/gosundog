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

// func BenchmarkAdd(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_ = v3.V{1, 2, 3}
// 	}
// }
