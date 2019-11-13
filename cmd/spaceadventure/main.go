package main

import "github.com/NikitaRubocki/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Solar System", Planets: []spaceadventure.Planet{
				spaceadventure.Planet{"Tattonine", "a desert planet"},
			},
		},
	)
}

