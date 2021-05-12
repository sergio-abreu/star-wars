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
	aPlanet, err := a.planetsRepository.GetById(query.ID)
	if err == nil {
		return aPlanet, nil
	}
	if err != planets.ErrPlanetNotFound {
		return planets.Planet{}, errors.Wrap(err, "failed to get planet by id")
	}
	aPlanet, err = a.planetsRepository.GetByName(query.Name)
	if err != nil {
		return planets.Planet{}, errors.Wrap(err, "failed to get planet by name")
	}
	return aPlanet, nil
}

func ErrNotFound(err error) bool {
	return errors.Cause(err) == planets.ErrPlanetNotFound
}
