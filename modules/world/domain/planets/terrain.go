package planets

import (
	"strings"
)

type Terrain string

func createTerrains(rawTerrains string) ([]Terrain, error) {
	var terrains []Terrain
	for _, rawTerrain := range strings.Split(rawTerrains, ",") {
		rawTerrain := strings.TrimSpace(rawTerrain)
		if len(rawTerrain) == 0 {
			rawTerrain = unknownTerrain
		}
		aTerrain, ok := allTerrains[rawTerrain]
		if !ok {
			return nil, NewErrTerrainNotFound(rawTerrain)
		}
		terrains = append(terrains, aTerrain)
	}
	return terrains, nil
}

const unknownTerrain = "unknown"

var allTerrains = map[string]Terrain{
	"airless asteroid": "airless asteroid",
	"barren":           "barren",
	"ash":              "ash",
	"caves":            "caves",
	"savannahs":        "savannahs",
	"savannas":         "savannas",
	"seas":             "seas",
	"cities":           "cities",
	"cityscape":        "cityscape",
	"desert":           "desert",
	"deserts":          "deserts",
	"forests":          "forests",
	"fungus forests":   "fungus forests",
	"gas giant":        "gas giant",
	"glaciers":         "glaciers",
	"ice canyons":      "ice canyons",
	"grass":            "grass",
	"grasslands":       "grasslands",
	"grassy hills":     "grassy hills",
	"rivers":           "rivers",
	"jungle":           "jungle",
	"jungles":          "jungles",
	"islands":          "islands",
	"lakes":            "lakes",
	"fields":           "fields",
	"rock arches":      "rock arches",
	"rocky deserts":    "rocky deserts",
	"mountain":         "mountain",
	"mountains":        "mountains",
	"valleys":          "valleys",
	"ocean":            "ocean",
	"oceans":           "oceans",
	"plateaus":         "plateaus",
	"reefs":            "reefs",
	"plains":           "plains",
	"hills":            "hills",
	"mesas":            "mesas",
	"cliffs":           "cliffs",
	"canyons":          "canyons",
	"rainforests":      "rainforests",
	"rock":             "rock",
	"rocky":            "rocky",
	"acid pools":       "acid pools",
	"rocky canyons":    "rocky canyons",
	"rocky islands":    "rocky islands",
	"sinkholes":        "sinkholes",
	"scrublands":       "scrublands",
	"swamp":            "swamp",
	"swamps":           "swamps",
	"toxic cloudsea":   "toxic cloudsea",
	"ice caves":        "ice caves",
	"mountain ranges":  "mountain ranges",
	"tundra":           "tundra",
	unknownTerrain:     unknownTerrain,
	"urban":            "urban",
	"bogs":             "bogs",
	"vines":            "vines",
	"verdant":          "verdant",
	"lava rivers":      "lava rivers",
	"volcanoes":        "volcanoes",
}
