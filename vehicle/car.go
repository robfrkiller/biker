package vehicle

import (
	"fmt"
)

const CarAcceleration = 7

type Car StandardCar

func (s *Car) Start() {
	if s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = true
	s.Engine.SetAcceleration(CarAcceleration)
	s.SpeedHistory = map[int]int{0: s.Speed}

	fmt.Println("啟動了轎車引擎")
}

func (s *Car) Stop() {
	if !s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = false
}

func (s *Car) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}

func (s *Car) IncreaseSpeed(time int) {
	if !s.Engine.IsStart {
		return
	}

	s.Speed += CarAcceleration
	s.SpeedHistory[time] = s.Speed
}

func (s *Car) CurrentSpeed() {
	fmt.Printf("轎車最後速度%d\n", s.Speed)
}
