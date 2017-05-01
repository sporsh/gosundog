package v3

import "math"

// V is a three dimensional vector
type V [3]float64

var (
	ZERO = V{0, 0, 0}
	X    = V{1, 0, 0}
	Y    = V{0, 1, 0}
	Z    = V{0, 0, 1}
)

// Add adds the three dimensional vector b to a
func (a *V) Add(b *V) *V {
	a[0] += b[0]
	a[1] += b[1]
	a[2] += b[2]
	return a
}

func (a *V) Dot(b *V) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func (a *V) Hadamard(b *V) *V {
	a[0] *= b[0]
	a[1] *= b[1]
	a[2] *= b[2]
	return a
}

func (v *V) Scale(f float64) *V {
	v[0] *= f
	v[1] *= f
	v[2] *= f
	return v
}

func (v *V) Len() float64 {
	return math.Sqrt(v.Len2())
}

func (v *V) Len2() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v *V) Normalize() *V {
	return v.Scale(1 / v.Len())
}

func (a *V) Sub(b *V) *V {
	a[0] -= b[0]
	a[1] -= b[1]
	a[2] -= b[2]
	return a
}

// Add computes the vector sum of two three dimensional vectors
func Add(a, b V) V {
	return *a.Add(&b)
}

// Cross computes the vector cross product of two three dimensional vectors
func Cross(a, b V) V {
	return V{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

// Dot computes the dod product of two three dimensional vectors
func Dot(a, b *V) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Hadamard returns the component wise product of two three dimensional vectors
func Hadamard(a, b V) V {
	return *a.Hadamard(&b)
}

// Len computes the length (magnitude) of a three dimensional vector
func Len(v V) float64 {
	return math.Sqrt(Len2(v))
}

// Len2 computes the squared length (magnitude) of a three dimensional vector
func Len2(v V) float64 {
	return Dot(&v, &v)
}

// Negate returns a three dimensional vector in the opposite direction
func Negate(v V) V {
	v[0], v[1], v[2] = -v[0], -v[1], -v[2]
	return v
}

// Normalize computes the unit length three dimensional vector
func Normalize(v V) V {
	return Scale(v, 1/Len(v))
}

// Scale comutes a three dimensional vector scaled by a factor
func Scale(v V, f float64) V {
	return *v.Scale(f)
}

// Sub computes the vector difference between two three dimensional vectors
func Sub(a, b V) V {
	return *a.Sub(&b)
}
