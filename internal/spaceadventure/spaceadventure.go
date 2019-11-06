package spaceadventure

import "fmt"

// PrintWelcome prints a welcome
func PrintWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

// GetResponseToPrompt takes a prompt and returns a response
func GetResponseToPrompt(prompt string) (response string) {
	fmt.Println(prompt)
    fmt.Scan(&response)
    return 
}

// PrintGreeting takes a name and prints a greeting
func PrintGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

// Travel asks whether a random planet should be chosen, or if the user should pick
func Travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = GetResponseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			TravelToRandomPlanet()
		} else if choice == "N" {
			TravelToPlanet(GetResponseToPrompt("Name the planet you would like to visit."))
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

// TravelToRandomPlanet prints a random planet
func TravelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

// TravelToPlanet prints the selected planet
func TravelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
