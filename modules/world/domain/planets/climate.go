package planets

import (
	"strings"
)

type Climate string

func (c Climate) ToString() string {
	return string(c)
}

type Climates []Climate

func (c Climates) ToStringSlice() []string {
	var stringSlice []string
	for _, climate := range c {
		stringSlice = append(stringSlice, climate.ToString())
	}
	return stringSlice
}

func createClimates(rawClimates string) ([]Climate, error) {
	var climates []Climate
	for _, rawClimate := range strings.Split(rawClimates, ",") {
		rawClimate := strings.TrimSpace(rawClimate)
		if len(rawClimate) == 0 {
			rawClimate = unknownClimate
		}
		aClimate, ok := allClimates[rawClimate]
		if !ok {
			return nil, NewErrClimateNotFound(rawClimate)
		}
		climates = append(climates, aClimate)
	}
	return climates, nil
}

const unknownClimate = "unknown"

var allClimates = map[string]Climate{
	"arid":                 "arid",
	"rocky":                "rocky",
	"windy":                "windy",
	"artificial temperate": "artificial temperate",
	"frigid":               "frigid",
	"frozen":               "frozen",
	"humid":                "humid",
	"murky":                "murky",
	"polluted":             "polluted",
	"superheated":          "superheated",
	"temperate":            "temperate",
	"subarctic":            "subarctic",
	"subartic":             "subartic",
	"hot":                  "hot",
	"artic":                "artic",
	"arctic":               "arctic",
	"moist":                "moist",
	"tropical":             "tropical",
	unknownClimate:         unknownClimate,
}
