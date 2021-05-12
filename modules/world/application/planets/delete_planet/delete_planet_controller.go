package delete_planet

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/domain/planets"
)

func NewDeletePlanetController(planetsRepository planets.Repository) *DeletePlanetController {
	return &DeletePlanetController{planetsRepository: planetsRepository}
}

type DeletePlanetController struct {
	planetsRepository planets.Repository
}

func (a *DeletePlanetController) DeletePlanet(query DeletePlanetCommand) error {
	err := a.planetsRepository.Delete(query.ID)
	return errors.Wrap(err, "failed to delete planet")
}