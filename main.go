package main

import (
	"biker/vehicle"
)

func main() {
	var v = &vehicle.Scooter{}

	v.Start()
	v.Acceleration()
}
