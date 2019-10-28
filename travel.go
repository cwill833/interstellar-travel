package main

import "fmt"

func fuelGuage(fuel int){
	fmt.Printf("You have %d gallons left in the tank", fuel)
}

func calculateFuel(planet string) int{
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

func greetPlanet(planet string){
	fmt.Printf("Welcome to %v \n", planet)
}

func cantFly(){
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int{
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost{
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main(){
	fuel := 1000000
	planetChoice := "Venus"

	fuel = flyToPlanet(planetChoice, fuel)

	fuelGuage(fuel)

}