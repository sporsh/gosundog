package v3

import "math"

// V is a three dimensional vector
type V struct {
	X, Y, Z float64
}

var (
	ZERO = V{0, 0, 0}
	X    = V{1, 0, 0}
	Y    = V{0, 1, 0}
	Z    = V{0, 0, 1}
)


// Add computes the vector sum of two three dimensional vectors
func Add(a, b V) V {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
	return a
}

// Cross computes the vector cross product of two three dimensional vectors
func Cross(a, b V) V {
	return V{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

// Dot computes the dod product of two three dimensional vectors
func Dot(a, b V) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Hadamard returns the component wise product of two three dimensional vectors
func Hadamard(a, b V) V {
	a.X *= b.X
	a.Y *= b.Y
	a.Z *= b.Z
	return a
}

// Len computes the length (magnitude) of a three dimensional vector
func Len(v V) float64 {
	return math.Sqrt(Len2(v))
}

// Len2 computes the squared length (magnitude) of a three dimensional vector
func Len2(v V) float64 {
	return Dot(v, v)
}

// Negate returns a three dimensional vector in the opposite direction
func Negate(v V) V {
	v.X, v.Y, v.Z = -v.X, -v.Y, -v.Z
	return v
}

// Normalize computes the unit length three dimensional vector
func Normalize(v V) V {
	return Scale(v, 1/Len(v))
}

// Scale comutes a three dimensional vector scaled by a factor
func Scale(v V, f float64) V {
	v.X *= f
	v.Y *= f
	v.Z *= f
	return v
}

// Sub computes the vector difference between two three dimensional vectors
func Sub(a, b V) V {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
	return a
}
