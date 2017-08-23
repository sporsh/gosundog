package sundog

import (
	"math"
	"math/rand"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type Camera struct {
	Origin      v3.V
	Basis       geometry.Basis
	Aperture    float64
	FocalLength float64
	FieldOfView float64
	AspectRatio float64
}

// RayThrough computes a ray in the world vector space
// that corresponds to the u, v screen coordinates between [-1, 1]
// as seen in the visible field from of cacmera
func (c *Camera) RayThrough(u, v, epsilon float64) geometry.Ray {
	origin := c.RandomOriginWithinAperture()
	target := v3.Add(
		c.Origin,
		c.Basis.ToLocal(v3.V{
			X: u * c.FieldOfView * c.FocalLength,
			Y: v * c.FieldOfView * c.FocalLength,
			Z: c.FocalLength,
		}),
	)
	direction := v3.Normalize(v3.Sub(target, origin))

	return geometry.Ray{
		Direction: direction,
		Origin:    origin,
		TMin:      epsilon,
		TMax:      math.Inf(1),
	}
}

// LookAt reorients the camera's basis so that it is looking straight at a target
func (c *Camera) LookAt(target v3.V) {

}

// RandomOriginWithinAperture returns, as a three dimensional vector, a point
// in world space that lies within the camera's aperture
func (c *Camera) RandomOriginWithinAperture() v3.V {
	if c.Aperture == 0 {
		return c.Origin
	}
	u1, u2 := rand.Float64(), rand.Float64()
	r := u2 * c.Aperture
	theta := u1 * 2 * math.Pi
	localPoint := v3.V{
		X: math.Cos(theta) * r,
		Y: math.Sin(theta) * r,
		Z: 0,
	}
	return v3.Add(
		c.Basis.ToLocal(localPoint),
		c.Origin,
	)
}
