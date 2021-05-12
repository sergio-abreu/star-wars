package planets

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestPlanet(t *testing.T) {
	t.Run("Do not create planet when name is empty", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := CreatePlanet("1", "", "", "")

		g.Expect(err).Should(
			MatchError(ErrEmptyPlanetName))
	})
	t.Run("Create planet with unknown climate and terrain when climate and terrain are empty", func(t *testing.T) {
		g := NewGomegaWithT(t)

		sut, err := CreatePlanet("1", "earth", "", "")

		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(sut.Name).Should(
			Equal("earth"))
		g.Expect(sut.Climates).Should(
			ConsistOf([]Climate{unknownClimate}))
		g.Expect(sut.Terrains).Should(
			ConsistOf([]Terrain{unknownTerrain}))
	})
	t.Run("Do not create planet when climate is not mapped", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := CreatePlanet("1", "earth", "no climate", "")

		g.Expect(err).Should(
			MatchError(`climate "no climate" not found`))
	})
	t.Run("Do not create planet when terrain is not mapped", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := CreatePlanet("1", "earth", "", "no terrain")

		g.Expect(err).Should(
			MatchError(`terrain "no terrain" not found`))
	})
	t.Run("Create planet when there are more than 1 climate and terrain", func(t *testing.T) {
		g := NewGomegaWithT(t)

		sut, err := CreatePlanet("1", "earth", "arid, rocky", "barren, ash")

		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(sut.Name).Should(
			Equal("earth"))
		g.Expect(sut.Climates).Should(
			ConsistOf([]Climate{"arid", "rocky"}))
		g.Expect(sut.Terrains).Should(
			ConsistOf([]Terrain{"barren", "ash"}))
	})
}
