package vehicle

import (
	"fmt"
)

type Truck struct {
	Engine
}

func (s *Truck) Start() {
	s.Engine.SetAcceleration(3)

	fmt.Println("啟動了卡車引擎")
}

func (s *Truck) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}
