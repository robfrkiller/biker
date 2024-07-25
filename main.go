package main

import (
	"biker/vehicle"
	"fmt"
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

	sMap := sync.Map{}

	ch := make(chan vehicle.VehicleSpeed, 1)
	wg := new(sync.WaitGroup)
	wg.Add(len(vehicles))

	go setHistory(&sMap, ch)

	for _, val := range vehicles {
		go startTesting(val, wg, ch)
	}

	wg.Wait()
	close(ch)

	for _, val := range vehicles {
		val.CurrentSpeed()
	}

	time.Sleep(time.Second)

	sMap.Range(func(key interface{}, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func bootup(v vehicle.Vehicle) {
	v.Start()
	v.Acceleration()
}

func startTesting(v vehicle.Vehicle, wg *sync.WaitGroup, ch chan<- vehicle.VehicleSpeed) {
	defer wg.Done()

	bootup(v)

	second := 0
	for range time.Tick(time.Second) {
		second++
		speed, err := v.IncreaseSpeed(second)
		if err == nil {
			ch <- vehicle.VehicleSpeed{
				Name:  v.GetName(),
				Speed: speed,
			}
		}

		if second == 5 {
			break
		}
	}

	v.Stop()
}

func setHistory(sMap *sync.Map, ch <-chan vehicle.VehicleSpeed) {
	for {
		if n, ok := <-ch; ok {
			(*sMap).Store(n.Name, n.Speed)
			fmt.Printf("目前%s車速：%d\n", n.Name, n.Speed)
		} else {
			break
		}
	}
}
