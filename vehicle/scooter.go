package vehicle

import (
	"fmt"
)

type Scooter struct {
	Engine
}

func (s *Scooter) Start() {
	s.Engine.SetAcceleration(1)

	fmt.Println("啟動了機車引擎")
}

func (s *Scooter) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}
