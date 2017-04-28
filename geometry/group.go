package geometry

type Group []Intersectable

func (g Group) Intersect(r Ray, epsilon float64) (closest Intersection, anyOk bool) {
	for _, intersectable := range g {
		if i, ok := intersectable.Intersect(r, epsilon); ok {
			if !anyOk {
				closest = i
			} else if i.T < closest.T {
				closest = i
			}
		}
	}
	return
}
