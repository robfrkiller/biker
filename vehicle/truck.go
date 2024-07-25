package vehicle

import (
	"errors"
	"fmt"
)

const TruckAcceleration = 3

type Truck StandardCar

func (s *Truck) GetName() string {
	return "卡車"
}

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

func (s *Truck) IncreaseSpeed(time int) (int, error) {
	if !s.Engine.IsStart {
		return 0, errors.New("engine must be running")
	}

	s.Speed += TruckAcceleration
	s.SpeedHistory[time] = s.Speed

	return s.Speed, nil
}

func (s *Truck) CurrentSpeed() {
	fmt.Printf("卡車最後速度%d\n", s.Speed)
}
