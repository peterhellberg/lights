package main

import (
	"testing"
	"time"
)

func TestStateSetCircadianValues(t *testing.T) {
	s := &State{}

	for _, tt := range []struct {
		value           string
		keyBrightness   int
		keyTemperature  int
		fillBrightness  int
		fillTemperature int
	}{
		{"1:02PM", 100, 5400, 75, 4900},
		{"3:04PM", 100, 4600, 75, 4100},
		{"6:35PM", 75, 2900, 50, 2400},
	} {
		t.Run(tt.value, func(t *testing.T) {
			now, err := time.Parse(time.Kitchen, tt.value)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			s.setCircadianValues(now)

			if got, want := s.Key.Brightness.number, tt.keyBrightness; got != want {
				t.Fatalf("s.Key.Brightness.number = %d, want %d", got, want)
			}

			if got, want := s.Key.Temperature.number, tt.keyTemperature; got != want {
				t.Fatalf("s.Key.Temperature.number = %d, want %d", got, want)
			}

			if got, want := s.Fill.Brightness.number, tt.fillBrightness; got != want {
				t.Fatalf("s.Fill.Brightness.number = %d, want %d", got, want)
			}

			if got, want := s.Fill.Temperature.number, tt.fillTemperature; got != want {
				t.Fatalf("s.Fill.Temperature.number = %d, want %d", got, want)
			}
		})
	}
}
