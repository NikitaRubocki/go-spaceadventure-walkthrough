package main

import "github.com/NikitaRubocki/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Solar System", Planets: mockPlanets(),
		},
	)
}


func mockPlanets() []spaceadventure.Planet {
	return []spaceadventure.Planet{
		spaceadventure.Planet{"Tatooine", "Desert planet"},
		spaceadventure.Planet{"Hoth", "Icy planet"},
		spaceadventure.Planet{Name:"Dagobah", Description:"Swamp planet"},
		spaceadventure.Planet{"Jaku", "Party Planet"},
		spaceadventure.Planet{"Vulcan", "Vulcan home planet"},
		spaceadventure.Planet{"Andoria", "Icy M-class planet"},
		spaceadventure.Planet{"Risa", "Tropical planet"},
		spaceadventure.Planet{"Shore Leave", "Amusement park planet"},
		spaceadventure.Planet{"Nibiru", "Jungle planet"},
		spaceadventure.Planet{"Delvia", "Plant people planet"}}
}
