package vehicle

import (
	"testing"
)

func TestScooter_Start(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		s := &Scooter{}
		s.Start()
		s.Stop()
		s.Start()

		if s.Speed != 0 {
			t.Errorf("s.Speed = %v, want %v", s.Speed, 0)
		}
		if s.Engine.IsStart != true {
			t.Errorf("s.Engine.IsStart = %v, want %v", s.Engine.IsStart, true)
		}
		if s.Engine.Acceleration != ScooterAcceleration {
			t.Errorf("s.Engine.Acceleration = %v, want %v", s.Engine.Acceleration, ScooterAcceleration)
		}
	})
	t.Run("test2", func(t *testing.T) {
		s := &Scooter{}
		s.Start()

		if s.Speed != 0 {
			t.Errorf("s.Speed = %v, want %v", s.Speed, 0)
		}
		if s.Engine.IsStart != true {
			t.Errorf("s.Engine.IsStart = %v, want %v", s.Engine.IsStart, true)
		}
		if s.Engine.Acceleration != ScooterAcceleration {
			t.Errorf("s.Engine.Acceleration = %v, want %v", s.Engine.Acceleration, ScooterAcceleration)
		}
	})
}

func TestScooter_Stop(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		s := &Scooter{}
		s.Stop()
		if s.Engine.IsStart != false {
			t.Errorf("s.Engine.IsStart = %v, want %v", s.Engine.IsStart, false)
		}
	})

	t.Run("test2", func(t *testing.T) {
		s := &Scooter{}
		s.Start()
		s.Stop()
		if s.Engine.IsStart != false {
			t.Errorf("s.Engine.IsStart = %v, want %v", s.Engine.IsStart, false)
		}
	})
}

func TestScooter_IncreaseSpeed(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		s := &Scooter{}
		s.IncreaseSpeed(1)
		if s.Speed != 0 {
			t.Errorf("s.Speed = %v, want %v", s.Speed, 0)
		}
	})
	t.Run("test2", func(t *testing.T) {
		s := &Scooter{}
		s.Start()
		s.IncreaseSpeed(1)
		if s.Speed != 1 {
			t.Errorf("s.Speed = %v, want %v", s.Speed, 1)
		}
		if s, ok := s.SpeedHistory[1]; !(ok && s == 1) {
			t.Errorf("s.SpeedHistory[1] = %v, want %v", s, 1)
		}
	})
}
