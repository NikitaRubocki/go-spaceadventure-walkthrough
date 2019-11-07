package main

import "github.com/NikitaRubocki/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	var ps = spaceadventure.PlanetarySystem{"Solar System"}
	spaceadventure.Start(ps)
}

