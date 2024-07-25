package vehicle

import "fmt"

type Vehicle interface {
	Start()
	Stop()
	Acceleration()
	IncreaseSpeed(time int)
	CurrentSpeed()
}

type StandardCar struct {
	Engine
	Speed        int
	SpeedHistory map[int]int
}

type Engine struct {
	IsStart      bool
	Acceleration int
}

func (e *Engine) SetAcceleration(n int) {
	e.Acceleration = n
}

func (s *Engine) String() string {
	return fmt.Sprintf("%d", s.Acceleration)
}
