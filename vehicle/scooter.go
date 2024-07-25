package vehicle

import (
	"fmt"
)

const ScooterAcceleration = 1

type Scooter StandardCar

func (s *Scooter) Start() {
	if s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = true
	s.Engine.SetAcceleration(ScooterAcceleration)
	s.SpeedHistory = map[int]int{0: s.Speed}

	fmt.Println("啟動了機車引擎")
}

func (s *Scooter) Stop() {
	if !s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = false
}

func (s *Scooter) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}

func (s *Scooter) IncreaseSpeed(time int) {
	if !s.Engine.IsStart {
		return
	}

	s.Speed += ScooterAcceleration
	s.SpeedHistory[time] = s.Speed
}

func (s *Scooter) CurrentSpeed() {
	fmt.Printf("機車最後速度%d\n", s.Speed)
}
