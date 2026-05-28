package bodies

type Geodesic struct {
	Ray *Ray
}

func (g *Geodesic) evalDirectionRadius() float64 {
	return g.Ray.radius + g.Ray.directionPhi*g.Ray.directionPhi - (c*c*g.Ray.speed)/(2*g.Ray.radius*g.Ray.radius)
}

func (g *Geodesic) evalDirectionPhi() float64 {
	return -2 * g.Ray.directionRadius * g.Ray.directionPhi / g.Ray.radius
}
