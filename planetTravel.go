package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("You will have:", fuel, "gallons of fuel left")
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Println("Welcome to", planet, "! We hope you enjoy your stay!")
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fy there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	fuel := 1000000
	var planetChoice string
	fmt.Println("Which planet would you like to fly to?(Mars, Venus, or Mercury)?: ")
	fmt.Scan(&planetChoice)
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
