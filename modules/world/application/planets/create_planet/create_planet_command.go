package create_planet

type CreatePlanetCommand struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Climates string `json:"climates"`
	Terrains string `json:"terrains"`
}
