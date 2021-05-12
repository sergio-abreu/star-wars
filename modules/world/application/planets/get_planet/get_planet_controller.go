package get_planet

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/domain/planets"
)

func NewGetPlanetController(planetsRepository planets.Repository) *GetPlanetController {
	return &GetPlanetController{planetsRepository: planetsRepository}
}

type GetPlanetController struct {
	planetsRepository planets.Repository
}

func (a *GetPlanetController) GetPlanet(query GetPlanetQuery) (planets.Planet, error) {
	planetId, err := planets.CreatePlanetID(query.ID)
	if err == nil {
		aPlanet, err := a.planetsRepository.GetById(planetId)
		if err != nil {
			return planets.Planet{}, errors.Wrap(err, "failed to get planet by id")
		}
		return aPlanet, nil
	}
	aPlanet, err := a.planetsRepository.GetByName(query.Name)
	if err != nil {
		return planets.Planet{}, errors.Wrap(err, "failed to get planet by name")
	}
	return aPlanet, nil
}

func ErrNotFound(err error) bool {
	return errors.Cause(err) == planets.ErrPlanetNotFound
}
