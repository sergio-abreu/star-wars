package planets

import "github.com/pkg/errors"

type Repository interface {
	GetById(id string) (Planet, error)
	GetByName(name string) (Planet, error)
	GetAll() ([]Planet, error)
	Delete(id string) error
	Save(aPlanet Planet) error
}

var ErrPlanetNotFound = errors.New("planet not found")
