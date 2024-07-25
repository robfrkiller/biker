package vehicle

import (
	"errors"
	"fmt"
)

const ScooterAcceleration = 1

type Scooter StandardCar

func (s *Scooter) GetName() string {
	return "機車"
}

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

func (s *Scooter) IncreaseSpeed(time int) (int, error) {
	if !s.Engine.IsStart {
		return 0, errors.New("engine must be running")
	}

	s.Speed += ScooterAcceleration
	s.SpeedHistory[time] = s.Speed

	return s.Speed, nil
}

func (s *Scooter) CurrentSpeed() {
	fmt.Printf("機車最後速度%d\n", s.Speed)
}
