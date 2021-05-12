package application

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	"github.com/sergio-vaz-abreu/star-wars/infrastructure/postgres"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestApplication(t *testing.T) {
	g := NewGomegaWithT(t)
	app, err := Load(GetConfig())
	g.Expect(err).Should(
		Not(HaveOccurred()))
	appErr := app.Run()
	g.Consistently(appErr).Should(
		Not(Receive()))
	guid, err := uuid.NewUUID()
	g.Expect(err).Should(
		Not(HaveOccurred()))
	t.Run("Create a planet", func(t *testing.T) {
		data := fmt.Sprintf(`{"id":"%s","name":"Yavin IV"}`, guid.String())
		httpResponse, err := http.Post("http://127.0.0.1:50059/api/v1/world/planets", "application/json", strings.NewReader(data))
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(httpResponse).Should(
			HaveHTTPStatus(http.StatusCreated))
		g.Expect(io.ReadAll(httpResponse.Body)).Should(
			MatchJSON(fmt.Sprintf(`{"id":"%s","name":"Yavin IV","climates":["unknown"],"terrains":["unknown"]}`, guid)))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
	t.Run("Get a planet by id", func(t *testing.T) {
		httpResponse, err := http.Get(fmt.Sprintf("http://127.0.0.1:50059/api/v1/world/planets/%s", guid))
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(httpResponse).Should(
			HaveHTTPStatus(http.StatusOK))
		g.Expect(io.ReadAll(httpResponse.Body)).Should(
			MatchJSON(fmt.Sprintf(`{"id":"%s","name":"Yavin IV","climates":["unknown"],"terrains":["unknown"]}`, guid)))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
	t.Run("Get a planet by name", func(t *testing.T) {
		httpResponse, err := http.Get("http://127.0.0.1:50059/api/v1/world/planets/Yavin IV")
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(httpResponse).Should(
			HaveHTTPStatus(http.StatusOK))
		g.Expect(io.ReadAll(httpResponse.Body)).Should(
			MatchJSON(fmt.Sprintf(`{"id":"%s","name":"Yavin IV","climates":["unknown"],"terrains":["unknown"]}`, guid)))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
	t.Run("Get all planets", func(t *testing.T) {
		httpResponse, err := http.Get("http://127.0.0.1:50059/api/v1/world/planets")
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(httpResponse).Should(
			HaveHTTPStatus(http.StatusOK))
		g.Expect(io.ReadAll(httpResponse.Body)).Should(
			ContainSubstring(fmt.Sprintf(`{"id":"%s","name":"Yavin IV","climates":["unknown"],"terrains":["unknown"]}`, guid)))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
	t.Run("Get a planet with movies apparitions", func(t *testing.T) {
		httpResponse, err := http.Get(fmt.Sprintf("http://127.0.0.1:50059/api/v1/fandom/planets/%s", guid))
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(httpResponse).Should(
			HaveHTTPStatus(http.StatusOK))
		g.Expect(io.ReadAll(httpResponse.Body)).Should(
			MatchJSON(fmt.Sprintf(`{"id":"%s","name":"Yavin IV","climates":["unknown"],"terrains":["unknown"],"movies":["A New Hope"]}`, guid)))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
	t.Run("Delete a planet", func(t *testing.T) {
		httpRequest, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://127.0.0.1:50059/api/v1/world/planets/%s", guid), nil)
		g.Expect(err).Should(
			Not(HaveOccurred()))
		httpResponse, err := http.DefaultClient.Do(httpRequest)
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(httpResponse).Should(
			HaveHTTPStatus(http.StatusNoContent))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
	t.Run("Shutting down app", func(t *testing.T) {
		app.Shutdown()
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
}

func GetConfig() Config {
	return Config{
		WebServerAddr: ":50059",
		SWApiBaseUrl:  "http://swapi.dev",
		Postgres: postgres.Config{
			User:     "starwars",
			Password: "P@ssword",
			Host:     "127.0.0.1",
			Port:     15432,
			Database: "starwars",
		},
	}
}
