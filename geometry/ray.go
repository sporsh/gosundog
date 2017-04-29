package geometry

import "github.com/sporsh/gosundog/v3"

type Ray struct {
	Origin, Direction v3.V
	TMin, TMax        float64
}
