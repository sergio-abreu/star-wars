package get_planets

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/movies/application/apparitions/get_movies_apparitions"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/get_planet"
)

func NewGetPlanetsController(apparitionsCtrl *get_movies_apparitions.GetMoviesApparitionsController, getPlanetCtrl *get_planet.GetPlanetController) *GetPlanetsController {
	return &GetPlanetsController{getPlanetCtrl: getPlanetCtrl, apparitionsCtrl: apparitionsCtrl}
}

type GetPlanetsController struct {
	apparitionsCtrl *get_movies_apparitions.GetMoviesApparitionsController
	getPlanetCtrl   *get_planet.GetPlanetController
}

func (c *GetPlanetsController) GetPlanet(query get_planet.GetPlanetQuery) (aPlanet GetPlanetsResponse, err error) {
	aPlanet.Planet, err = c.getPlanetCtrl.GetPlanet(query)
	if err != nil {
		return aPlanet, errors.Wrap(err, "failed to get planet")
	}
	aPlanet.Movies, err = c.apparitionsCtrl.GetMoviesApparitions(get_movies_apparitions.GetMoviesApparitionsQuery{PlanetName: aPlanet.Name})
	if err != nil && !get_movies_apparitions.ErrNotFound(err) {
		return aPlanet, errors.Wrap(err, "failed to get movies apparitions")
	}
	return aPlanet, nil
}
