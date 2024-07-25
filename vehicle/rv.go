package vehicle

import (
	"fmt"
)

type RV struct {
	Engine
}

func (s *RV) Start() {
	s.Engine.SetAcceleration(5)

	fmt.Println("啟動了休旅車引擎")
}

func (s *RV) Acceleration() {
	fmt.Printf("加速度為%s\n", &s.Engine)
}
