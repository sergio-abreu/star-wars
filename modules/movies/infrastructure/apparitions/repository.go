package apparitions

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/movies/domain/apparitions"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"sync"
)

func NewWebRepository(baseUrl string) (*WebRepository, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}
	return &WebRepository{baseUrl: parsedUrl}, nil
}

type WebRepository struct {
	baseUrl *url.URL
}

func (r *WebRepository) GetMoviesByPlanetName(planet string) ([]apparitions.Movie, error) {
	r.baseUrl.Path = "/api/planets"
	baseUrl := r.baseUrl.String()
	allPlanetsUrls := []string{
		fmt.Sprintf("%s?page=%d", baseUrl, 1),
		fmt.Sprintf("%s?page=%d", baseUrl, 2),
		fmt.Sprintf("%s?page=%d", baseUrl, 3),
		fmt.Sprintf("%s?page=%d", baseUrl, 4),
		fmt.Sprintf("%s?page=%d", baseUrl, 5),
		fmt.Sprintf("%s?page=%d", baseUrl, 6),
	}
	aPlanet, err := r.getPlanetByName(planet, allPlanetsUrls...)
	if err != nil {
		return nil, err
	}
	return r.getMoviesBySwapiPlanet(aPlanet)
}

type swapiResponse struct {
	Next    string
	Planets []swapiPlanet `json:"results"`
}

type swapiPlanet struct {
	Name      string
	FilmsUrls []string `json:"films"`
}

type swapiMovie struct {
	Title string
}

func (r *WebRepository) getPlanetByName(name string, allPlanetsUrls ...string) (swapiPlanet, error) {
	aPlanetCh := make(chan swapiPlanet, len(allPlanetsUrls)*60)
	defer close(aPlanetCh)
	done := make(chan struct{}, 1)
	errCh := make(chan error, len(allPlanetsUrls))
	defer close(errCh)
	var wg sync.WaitGroup
	for _, allPlanetsUrl := range allPlanetsUrls {
		wg.Add(1)
		go func(allPlanetsUrl string) {
			defer wg.Done()
			var response swapiResponse
			err := r.sendGetRequest(allPlanetsUrl, &response)
			if err != nil {
				errCh <- errors.Wrap(err, "failed to get all planets")
				return
			}
			for _, aPlanet := range response.Planets {
				if name != aPlanet.Name {
					continue
				}
				aPlanetCh <- aPlanet
				return
			}
		}(allPlanetsUrl)
	}
	go func() {
		wg.Wait()
		done <- struct{}{}
		close(done)
	}()
	select {
	case aPlanet := <-aPlanetCh:
		return aPlanet, nil
	case err := <-errCh:
		return swapiPlanet{}, err
	case <-done:
		return swapiPlanet{}, apparitions.ErrMovieNotFound
	}
}

func (r *WebRepository) getMoviesBySwapiPlanet(aPlanet swapiPlanet) ([]apparitions.Movie, error) {
	movieCh := make(chan swapiMovie, len(aPlanet.FilmsUrls))
	var wg sync.WaitGroup
	for _, movieUrl := range aPlanet.FilmsUrls {
		wg.Add(1)
		go func(movieUrl string) {
			defer wg.Done()
			var movie swapiMovie
			err := r.sendGetRequest(movieUrl, &movie)
			if err != nil {
				logrus.WithError(err).Error("failed to get film")
				return
			}
			movieCh <- movie
		}(movieUrl)
	}
	wg.Wait()
	close(movieCh)
	movies := make([]apparitions.Movie, 0)
	for movie := range movieCh {
		movies = append(movies, apparitions.NewMovie(movie.Title))
	}
	return movies, nil
}

func (r *WebRepository) sendGetRequest(url string, data interface{}) error {
	httpResponse, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "failed to send http get request")
	}
	httpBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read http body")
	}
	if httpResponse.StatusCode != http.StatusOK {
		return errors.Errorf("%d - %s", httpResponse.StatusCode, httpBody)
	}
	err = json.Unmarshal(httpBody, &data)
	if err != nil {
		return errors.Wrap(err, "failed to json decode http body")
	}
	return nil
}
