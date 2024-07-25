package vehicle

import (
	"fmt"
)

type Car struct {
	Engine
}

func (s *Car) Start() {
	s.Engine.SetAcceleration(7)

	fmt.Println("啟動了轎車引擎")
}

func (s *Car) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}
