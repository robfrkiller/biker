package vehicle

import (
	"fmt"
)

const TruckAcceleration = 3

type Truck StandardCar

func (s *Truck) Start() {
	if s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = true
	s.Engine.SetAcceleration(TruckAcceleration)
	s.SpeedHistory = map[int]int{0: s.Speed}

	fmt.Println("啟動了卡車引擎")
}

func (s *Truck) Stop() {
	if !s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = false
}

func (s *Truck) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}

func (s *Truck) IncreaseSpeed(time int) {
	if !s.Engine.IsStart {
		return
	}

	s.Speed += TruckAcceleration
	s.SpeedHistory[time] = s.Speed
}

func (s *Truck) CurrentSpeed() {
	fmt.Printf("卡車最後速度%d\n", s.Speed)
}
