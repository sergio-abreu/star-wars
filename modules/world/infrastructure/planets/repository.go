package planets

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/domain/planets"
)

func NewSqlRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{db: db}
}

type SqlRepository struct {
	db *sql.DB
}

func (r *SqlRepository) GetById(id string) (planets.Planet, error) {
	query := "SELECT id, name, climates, terrains FROM planets where id = $1;"
	row := r.db.QueryRow(query, id)
	return r.scanPlanet(row)
}

func (r *SqlRepository) GetByName(name string) (planets.Planet, error) {
	query := "SELECT id, name, climates, terrains FROM planets where name = $1"
	row := r.db.QueryRow(query, name)
	return r.scanPlanet(row)
}

func (r *SqlRepository) GetAll() ([]planets.Planet, error) {
	query := "SELECT id, name, climates, terrains FROM planets;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query planets on database")
	}
	var allPlanets []planets.Planet
	for rows.Next() {
		var planetID, planetName, planetClimates, planetTerrains string
		err := rows.Scan(&planetID, &planetName, &planetClimates, &planetTerrains)
		if err != nil {
			return nil, errors.Wrap(err, "failed to query planets on database")
		}
		aPlanet, err := planets.CreatePlanet(planetID, planetName, planetClimates, planetTerrains)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create planet from database invalid state")
		}
		allPlanets = append(allPlanets, aPlanet)
	}
	return allPlanets, nil
}

func (r *SqlRepository) Delete(id string) error {
	query := "DELETE FROM planets WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	return errors.Wrap(err, "failed to delete planet from database")
}

func (r *SqlRepository) Save(aPlanet planets.Planet) error {
	query := "INSERT INTO planets (id, name, climates, terrains) VALUES ();"
	result, err := r.db.Exec(query, aPlanet.ID, aPlanet.Name, aPlanet.Climates, aPlanet.Terrains)
	if err != nil {
		return errors.Wrap(err, "failed to create planet into database")
	}
	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		return errors.Wrap(err, "failed to create planet into database: no rows affected")
	}
	return nil
}

func (r *SqlRepository) scanPlanet(row *sql.Row) (planets.Planet, error) {
	var planetID, planetName, planetClimates, planetTerrains string
	err := row.Scan(&planetID, &planetName, &planetClimates, &planetTerrains)
	if err != sql.ErrNoRows {
		return planets.Planet{}, planets.ErrPlanetNotFound
	}
	if err != nil {
		return planets.Planet{}, errors.Wrap(err, "failed to scan planet on database")
	}
	aPlanet, err := planets.CreatePlanet(planetID, planetName, planetClimates, planetTerrains)
	if err != nil {
		return planets.Planet{}, errors.Wrap(err, "failed to create planet from database invalid state")
	}
	return aPlanet, nil
}
