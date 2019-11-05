package main

import "fmt"

import "github.com/NikitaRubocki/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.printWelcome()
	spaceadventure.printGreeting(spaceadventure.getResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	spaceadventure.travel()
}

