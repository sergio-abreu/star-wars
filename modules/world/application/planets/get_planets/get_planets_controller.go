package get_planets

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/domain/planets"
)

func NewGetPlanetsController(planetsRepository planets.Repository) *GetPlanetsController {
	return &GetPlanetsController{planetsRepository: planetsRepository}
}

type GetPlanetsController struct {
	planetsRepository planets.Repository
}

func (a *GetPlanetsController) GetPlanets() ([]planets.Planet, error) {
	allPlanets, err := a.planetsRepository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get planets")
	}
	return allPlanets, nil
}
