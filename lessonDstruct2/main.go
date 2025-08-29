package main

import (
	"fmt"
)

func main() {
	// create an stand alone anonymous struct
	myCar := struct {
		brand string
		model string
	}{
		brand: "Toyota",
		model: "Camry",
	}

	fmt.Println("Anonymous struct:", myCar)
	fmt.Println("nestsed anonymous struct:", yourCar)
}

// a named struct containing an anonymous struct
type car struct {
	brand   string
	model   string
	doors   int
	mileage int
	// wheel is a field containing an anonymous function
	wheel struct {
		radius   int
		material string
	}
}

var yourCar = car{
	brand:   "Rezvani",
	model:   "Vengeance",
	doors:   4,
	mileage: 35000,
	wheel: struct {
		radius   int
		material string
	}{
		radius:   35,
		material: "alloy",
	},
}
