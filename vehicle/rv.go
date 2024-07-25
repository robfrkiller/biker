package vehicle

import (
	"errors"
	"fmt"
)

const RVAcceleration = 5

type RV StandardCar

func (s *RV) GetName() string {
	return "休旅車"
}

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

func (s *RV) IncreaseSpeed(time int) (int, error) {
	if !s.Engine.IsStart {
		return 0, errors.New("engine must be running")
	}

	s.Speed += RVAcceleration
	s.SpeedHistory[time] = s.Speed

	return s.Speed, nil
}

func (s *RV) CurrentSpeed() {
	fmt.Printf("休旅車最後速度%d\n", s.Speed)
}
