package main

import "fmt"

import "github.com/NikitaRubocki/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.PrintWelcome()
	spaceadventure.PrintGreeting(spaceadventure.GetResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	spaceadventure.Travel()
}

