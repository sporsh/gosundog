package geometry

type Bound struct {
	Bounds   Intersectable
	Geometry Intersectable
}

func (b *Bound) Intersect(r *Ray, epsilon float64) (i Intersection, ok bool) {
	if i, ok = b.Bounds.Intersect(r, epsilon); ok {
		return b.Geometry.Intersect(r, epsilon)
	}
	return
}
