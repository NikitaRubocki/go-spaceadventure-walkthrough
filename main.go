package main

import "fmt"

func main() {
	printWelcome()
	printGreeting(getResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	travel()
}

func printWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func getResponseToPrompt(prompt string) (response string) {
	fmt.Println(prompt)
    fmt.Scan(&response)
    return 
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getResponseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			travelToRandomPlanet()
		} else if choice == "N" {
			travelToPlanet(getResponseToPrompt("Name the planet you would like to visit."))
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}


func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
