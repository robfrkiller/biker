package main

import (
	"biker/vehicle"
)

func main() {
	bootup(&vehicle.Scooter{})
	bootup(&vehicle.Truck{})
	bootup(&vehicle.RV{})
	bootup(&vehicle.Car{})
}

func bootup(v vehicle.Vehicle) {
	v.Start()
	v.Acceleration()
}
