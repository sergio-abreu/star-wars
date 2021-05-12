package get_movies_apparitions

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/movies/domain/apparitions"
)

func NewGetMoviesApparitionsController(apparitionsRepository apparitions.Repository) *GetMoviesApparitionsController {
	return &GetMoviesApparitionsController{apparitionsRepository: apparitionsRepository}
}

type GetMoviesApparitionsController struct {
	apparitionsRepository apparitions.Repository
}

func (c *GetMoviesApparitionsController) GetMoviesApparitions(query GetMoviesApparitionsQuery) ([]apparitions.Movie, error) {
	return c.apparitionsRepository.GetMoviesByPlanetName(query.PlanetName)
}

func ErrNotFound(err error) bool {
	return errors.Cause(err) == apparitions.ErrMovieNotFound
}
