package vehicle

import (
	"fmt"
)

const RVAcceleration = 5

type RV StandardCar

func (s *RV) Start() {
	if s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = true
	s.Engine.SetAcceleration(RVAcceleration)
	s.SpeedHistory = map[int]int{0: s.Speed}

	fmt.Println("啟動了休旅車引擎")
}

func (s *RV) Stop() {
	if !s.Engine.IsStart {
		return
	}

	s.Engine.IsStart = false
}

func (s *RV) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}

func (s *RV) IncreaseSpeed(time int) {
	if !s.Engine.IsStart {
		return
	}

	s.Speed += RVAcceleration
	s.SpeedHistory[time] = s.Speed
}

func (s *RV) CurrentSpeed() {
	fmt.Printf("休旅車最後速度%d\n", s.Speed)
}
