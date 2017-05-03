package sundog

import (
	"math"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type Sampler interface {
	Sample(r geometry.Ray) Sample
}

type Sample struct {
	Radiance, Weight v3.V
}

func (s Sample) RGBA() (r, g, b, a uint32) {
	r = uint32(math.Min(1, math.Max(0, s.Radiance.X)) * 0xffff)
	g = uint32(math.Min(1, math.Max(0, s.Radiance.Y)) * 0xffff)
	b = uint32(math.Min(1, math.Max(0, s.Radiance.Z)) * 0xffff)
	a = 0xffff
	return
}
