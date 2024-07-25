package vehicle

import "fmt"

type Vehicle interface {
	Start()
	Acceleration()
}

type Engine struct {
	Acceleration int
}

func (e *Engine) SetAcceleration(n int) {
	e.Acceleration = n
}

func (s *Engine) String() string {
	return fmt.Sprintf("%d", s.Acceleration)
}
