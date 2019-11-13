package spaceadventure

// PlanetarySystem for creating systems of planets
type PlanetarySystem struct {
	Name string
	Planets []Planet
}

// NumberOfPlanets calculates planets in the planetary system
func (ps PlanetarySystem) NumberOfPlanets() int{
	return len(ps.Planets)
}
