package geometry

type Bound struct {
	Bounds   Intersectable
	Geometry Intersectable
}

func (b *Bound) Intersect(r *Ray) (i Intersection, ok bool) {
	if i, ok = b.Bounds.Intersect(r); ok {
		return b.Geometry.Intersect(r)
	}
	return
}
