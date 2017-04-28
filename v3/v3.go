package v3

import "math"

// V is a three dimensional vector
type V [3]float64

// Add computes the vector sum of two three dimensional vectors
func Add(a, b V) V {
	a[0] += b[0]
	a[1] += b[1]
	a[2] += b[2]
	return a
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
func Dot(a, b V) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Len computes the length (magnitude) of a three dimensional vector
func Len(v V) float64 {
	return math.Sqrt(Len2(v))
}

// Len2 computes the squared length (magnitude) of a three dimensional vector
func Len2(v V) float64 {
	return Dot(v, v)
}

// Normalize computes the unit length three dimensional vector
func Normalize(v V) V {
	return Scale(v, 1/Len(v))
}

// Scale comutes a three dimensional vector scaled by a factor
func Scale(v V, f float64) V {
	v[0] *= f
	v[1] *= f
	v[2] *= f
	return v
}

// Sub computes the vector difference between two three dimensional vectors
func Sub(a, b V) V {
	a[0] -= b[0]
	a[1] -= b[1]
	a[2] -= b[2]
	return a
}
