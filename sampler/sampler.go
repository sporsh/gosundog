package sampler

import (
	"math"

	"github.com/sporsh/gosundog/geometry"
	"github.com/sporsh/gosundog/v3"
)

type Sampler interface {
	Sample(r geometry.Ray) Sample
}

type Sample v3.V

func (s Sample) RGBA() (r, g, b, a uint32) {
	r = uint32(math.Min(1, math.Max(0, s[0])) * 0xffff)
	g = uint32(math.Min(1, math.Max(0, s[1])) * 0xffff)
	b = uint32(math.Min(1, math.Max(0, s[2])) * 0xffff)
	a = 0xffff
	return
}
