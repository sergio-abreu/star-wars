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
	planetId, err := planets.CreatePlanetID(query.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create planet id")
	}
	err = a.planetsRepository.Delete(planetId)
	return errors.Wrap(err, "failed to delete planet")
}

func ErrInvalidInput(err error) bool {
	cause := errors.Cause(err)
	return cause == planets.ErrInvalidPlanetID
}
