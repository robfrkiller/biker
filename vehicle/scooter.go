package vehicle

import (
	"fmt"
)

type Scooter struct {
	ScooterEngine
}

type ScooterEngine struct {
	Acceleration int
}

func (s *Scooter) Start() {
	s.ScooterEngine.Start()

	fmt.Println("啟動了機車引擎")
}

func (s *Scooter) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.ScooterEngine)
}

func (e *ScooterEngine) Start() {
	e.Acceleration = 1
}

func (s *ScooterEngine) String() string {
	return fmt.Sprintf("%d", s.Acceleration)
}
