package main

import (
	"biker/vehicle"
	"sync"
	"time"
)

func main() {
	vehicles := []vehicle.Vehicle{
		&vehicle.Scooter{},
		&vehicle.Truck{},
		&vehicle.RV{},
		&vehicle.Car{},
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(vehicles))

	for _, val := range vehicles {
		go startTesting(val, wg)
	}

	wg.Wait()

	for _, val := range vehicles {
		val.CurrentSpeed()
	}
}

func bootup(v vehicle.Vehicle) {
	v.Start()
	v.Acceleration()
}

func startTesting(v vehicle.Vehicle, wg *sync.WaitGroup) {
	defer wg.Done()

	bootup(v)

	second := 0
	for range time.Tick(time.Second) {
		second++
		v.IncreaseSpeed(second)

		if second == 5 {
			break
		}
	}

	v.Stop()
}
