package geometry

type Group []Intersectable

func (g Group) Intersect(r Ray, epsilon float64) (closest Intersection, anyOk bool) {
	for _, intersectable := range g {
		if i, ok := intersectable.Intersect(r, epsilon); ok {
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
