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
	r = uint32(math.Min(1, math.Max(0, s.Radiance[0])) * 0xffff)
	g = uint32(math.Min(1, math.Max(0, s.Radiance[1])) * 0xffff)
	b = uint32(math.Min(1, math.Max(0, s.Radiance[2])) * 0xffff)
	a = 0xffff
	return
}
