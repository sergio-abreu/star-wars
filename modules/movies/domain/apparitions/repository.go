package apparitions

import "errors"

var ErrMovieNotFound = errors.New("movie not found")

type Repository interface {
	GetMoviesByPlanetName(planet string) ([]Movie, error)
}
