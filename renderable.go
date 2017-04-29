package sundog

import (
	"github.com/sporsh/gosundog/geometry"
)

type Renderable struct {
	geometry.Intersectable
	Material Material
}
