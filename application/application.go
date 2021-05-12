package application

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/application/controllers"
	"github.com/sergio-vaz-abreu/star-wars/infrastructure/postgres"
	"github.com/sergio-vaz-abreu/star-wars/modules/fandom/application/get_planet_with_apparitions"
	"github.com/sergio-vaz-abreu/star-wars/modules/movies/application/apparitions/get_movies_apparitions"
	"github.com/sergio-vaz-abreu/star-wars/modules/movies/infrastructure/apparitions"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/create_planet"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/delete_planet"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/get_planet"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/get_planets"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/infrastructure/planets"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Load(config Config) (*Application, error) {
	apparitionRepository, err := apparitions.NewWebRepository(config.SWApiBaseUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create apparition web repository")
	}
	db, err := postgres.NewDatabase(config.Postgres)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create postgres database")
	}
	planetRepository := planets.NewSqlRepository(db)
	getMoviesApparitionsCtrl := get_movies_apparitions.NewGetMoviesApparitionsController(apparitionRepository)
	getPlanetCtrl := get_planet.NewGetPlanetController(planetRepository)
	engine := gin.Default()
	api := engine.Group("/api/v1")
	api.POST("/world/planets", controllers.V1CreatePlanet(create_planet.NewCreatePlanetController(planetRepository)))
	api.GET("/world/planets/:id", controllers.V1GetPlanet(getPlanetCtrl))
	api.GET("/world/planets", controllers.V1GetPlanets(get_planets.NewGetPlanetsController(planetRepository)))
	api.DELETE("/world/planets/:id", controllers.V1DeletePlanets(delete_planet.NewDeletePlanetController(planetRepository)))
	api.GET("/fandom/planets/:id", controllers.V1GetFandomPlanet(get_planet_with_apparitions.NewGetPlanetsController(getMoviesApparitionsCtrl, getPlanetCtrl)))
	return &Application{&http.Server{Handler: engine, Addr: config.WebServerAddr}}, nil
}

type Application struct {
	httpServer *http.Server
}

func (a *Application) Run() <-chan error {
	appErr := make(chan error)
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appErr <- errors.Wrap(err, "failed listening http server")
		}
	}()
	return appErr
}

func (a *Application) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.WithError(err).Error("failed shutting down http server")
	}
}
