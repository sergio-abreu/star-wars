package get_planets

import (
	"github.com/sergio-vaz-abreu/star-wars/modules/movies/domain/apparitions"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/domain/planets"
)

type GetPlanetsResponse struct {
	planets.Planet
	Movies []apparitions.Movie
}
