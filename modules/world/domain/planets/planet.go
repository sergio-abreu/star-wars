package planets

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrEmptyPlanetName = errors.New("empty planet name")
var ErrInvalidPlanetID = errors.New("invalid planet id format (uuid)")

func CreatePlanet(id, name, climate, terrain string) (aPlanet Planet, err error) {
	aPlanet.ID, err = CreatePlanetID(id)
	if err != nil {
		return
	}
	aPlanet.Name, err = createName(name)
	if err != nil {
		return
	}
	aPlanet.Climates, err = createClimates(climate)
	if err != nil {
		return
	}
	aPlanet.Terrains, err = createTerrains(terrain)
	if err != nil {
		return
	}
	return
}

func createName(name string) (string, error) {
	if len(name) == 0 {
		return name, ErrEmptyPlanetName
	}
	return name, nil
}

func CreatePlanetID(id string) (PlanetID, error) {
	planetId, err := uuid.Parse(id)
	if err != nil {
		return PlanetID{}, ErrInvalidPlanetID
	}
	return PlanetID{id: planetId}, nil
}

type PlanetID struct {
	id uuid.UUID
}

func (p PlanetID) String() string {
	return p.id.String()
}

func (p PlanetID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", p.String())), nil
}

type Planet struct {
	ID       PlanetID `json:"id"`
	Name     string   `json:"name"`
	Climates Climates `json:"climates"`
	Terrains Terrains `json:"terrains"`
}
