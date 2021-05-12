package planets

import "errors"

var ErrEmptyPlanetName = errors.New("empty planet name")

func CreatePlanet(id, name, climate, terrain string) (Planet, error) {
	planet := Planet{ID: id, Name: name}
	if len(planet.Name) == 0 {
		return planet, ErrEmptyPlanetName
	}
	var err error
	planet.Climates, err = createClimates(climate)
	if err != nil {
		return planet, err
	}
	planet.Terrains, err = createTerrains(terrain)
	if err != nil {
		return planet, err
	}
	return planet, nil
}

type Planet struct {
	ID       string
	Name     string
	Climates []Climate
	Terrains []Terrain
}
