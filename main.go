package main

import (
	"biker/vehicle"
)

func main() {
	vehicles := []vehicle.Vehicle{
		&vehicle.Scooter{},
		&vehicle.Truck{},
		&vehicle.RV{},
		&vehicle.Car{},
	}

	for _, val := range vehicles {
		bootup(val)
	}
}

func bootup(v vehicle.Vehicle) {
	v.Start()
	v.Acceleration()
}
