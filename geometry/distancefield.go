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

func (df *DistanceField) Intersect(r *Ray, epsilon float64) (i Intersection, ok bool) {
	var distance float64
	i.Point = r.Origin

	for i.T = r.TMin + epsilon; i.T < math.Min(r.TMax, 10); i.T += distance {
		// i.Point = v3.Add(r.Origin, v3.Scale(r.Direction, i.T))
		i.Point = v3.Add(i.Point, v3.Scale(r.Direction, i.T))
		distance = math.Abs(df.distance(i.Point))
		if distance < epsilon {
			// i.Geometry = df
			i.Basis = ArbritraryBasisForNormal(df.getNormal(i.Point, epsilon))
			// fmt.Println("GOT INTERSECTION", i)
			return i, true
		}
	}
	return i, false
}

func (df *DistanceField) getNormal(point v3.V, epsilon float64) v3.V {
	d := func(v v3.V) float64 {
		// point = v3.Add(point, v)
		// return df.distance(point)
		return df.distance(v3.Add(point, v))
	}

	return v3.Normalize(v3.V{
		X: d(v3.V{X: -epsilon, Y: 0, Z: 0}) - d(v3.V{X: epsilon, Y: 0, Z: 0}),
		Y: d(v3.V{X: 0, Y: -epsilon, Z: 0}) - d(v3.V{X: 0, Y: epsilon, Z: 0}),
		Z: d(v3.V{X: 0, Y: 0, Z: -epsilon}) - d(v3.V{X: 0, Y: 0, Z: epsilon}),
	})
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
