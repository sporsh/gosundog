package geometry

import (
	"math"

	"github.com/sporsh/gosundog/v3"
)

type DistanceFunction func(v3.V) float64

type DistanceField struct {
	distance DistanceFunction
}

func NewDistanceField(distance DistanceFunction) *DistanceField {
	return &DistanceField{
		distance: distance,
	}
}

func (df *DistanceField) Intersect(r *Ray) (i Intersection, ok bool) {
	var distance float64
	i.Point = r.Origin

	for i.T = r.TMin; i.T < math.Min(r.TMax, 10); i.T += distance {
		// i.Point = v3.Add(r.Origin, v3.Scale(r.Direction, i.T))
		i.Point = v3.Add(i.Point, v3.Scale(r.Direction, i.T))
		distance = (df.distance(i.Point))
		if distance <= 0 {
			// i.Geometry = df
			i.Basis = ArbritraryBasisForNormal(v3.Negate(df.getNormal(i.Point)))
			// fmt.Println("GOT INTERSECTION", i)
			return i, true
		}
	}
	return i, false
}

func (df *DistanceField) getNormal(point v3.V) v3.V {
	d := func(v v3.V) float64 {
		// point = v3.Add(point, v)
		// return df.distance(point)
		return df.distance(v3.Add(point, v))
	}

	return v3.Normalize(v3.V{
		X: d(v3.V{X: -0.0001, Y: 0, Z: 0}) - d(v3.V{X: 0.0001, Y: 0, Z: 0}),
		Y: d(v3.V{X: 0, Y: -0.0001, Z: 0}) - d(v3.V{X: 0, Y: 0.0001, Z: 0}),
		Z: d(v3.V{X: 0, Y: 0, Z: -0.0001}) - d(v3.V{X: 0, Y: 0, Z: 0.0001}),
	})
}

func NewDSphere(r float64) *DistanceField {
	return &DistanceField{
		distance: func(point v3.V) float64 {
			return v3.Len(point) - r
		},
	}
}

func NewTorus(major, minor float64) *DistanceField {
	return &DistanceField{
		distance: func(point v3.V) float64 {
			l0 := math.Sqrt(point.X*point.X+point.Y*point.Y) - major
			l1 := point.Z
			return math.Sqrt(l0*l0+l1*l1) - minor
		},
	}
}

func NewTorus88(major, minor float64) *DistanceField {
	return &DistanceField{
		distance: func(point v3.V) float64 {
			l0 := math.Sqrt(point.X*point.X+point.Y*point.Y) - major
			l1 := point.Z
			return math.Sqrt(l0*l0+l1*l1) - minor
		},
	}
}
