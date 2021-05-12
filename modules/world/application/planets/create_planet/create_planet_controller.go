package create_planet

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/domain/planets"
)

func NewCreatePlanetController(planetsRepository planets.Repository) *CreatePlanetController {
	return &CreatePlanetController{planetsRepository: planetsRepository}
}

type CreatePlanetController struct {
	planetsRepository planets.Repository
}

func (a *CreatePlanetController) CreatePlanet(command CreatePlanetCommand) (planets.Planet, error) {
	planetId, err := planets.CreatePlanetID(command.ID)
	if err != nil {
		return planets.Planet{}, errors.Wrap(err, "failed to create planet id")
	}
	aPlanet, err := a.planetsRepository.GetById(planetId)
	if err == nil {
		return aPlanet, nil
	}
	if err != planets.ErrPlanetNotFound {
		return planets.Planet{}, errors.Wrap(err, "failed to check if planet already exists")
	}
	aPlanet, err = planets.CreatePlanet(command.ID, command.Name, command.Climates, command.Terrains)
	if err != nil {
		return planets.Planet{}, errors.Wrap(err, "failed to create planet")
	}
	err = a.planetsRepository.Save(aPlanet)
	if err != nil {
		return planets.Planet{}, errors.Wrap(err, "failed to save planet into repository")
	}
	return aPlanet, nil
}

func ErrInvalidInput(err error) bool {
	cause := errors.Cause(err)
	isEmptyPlanetName := cause == planets.ErrEmptyPlanetName
	_, isClimateNotFound := cause.(planets.ErrClimateNotFound)
	_, isTerrainNotFound := cause.(planets.ErrTerrainNotFound)
	isInvalidPlanetID := cause == planets.ErrInvalidPlanetID
	return isEmptyPlanetName || isClimateNotFound || isTerrainNotFound || isInvalidPlanetID
}
