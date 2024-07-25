package vehicle

import (
	"testing"
)

func TestEngine_SetAcceleration(t *testing.T) {
	type fields struct {
		Acceleration int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "test1", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{}
			e.SetAcceleration(1)
			if e.Acceleration != tt.want {
				t.Errorf("SetAcceleration() = %v, want %v", e.Acceleration, tt.want)
			}
		})
	}
}

func TestEngine_String(t *testing.T) {
	type fields struct {
		Acceleration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "test1", want: "0"},
		{name: "test2", fields: fields{Acceleration: 2}, want: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Engine{
				Acceleration: tt.fields.Acceleration,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Engine.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
