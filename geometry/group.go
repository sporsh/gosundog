package geometry

type Group []Intersectable

func (g *Group) Intersect(r *Ray) (closest Intersection, anyOk bool) {
	for _, intersectable := range *g {
		if i, ok := intersectable.Intersect(r); ok {
			i.Geometry = intersectable
			if !anyOk {
				closest = i
				anyOk = true
			} else if i.T < closest.T {
				closest = i
			}
		}
	}
	return
}
