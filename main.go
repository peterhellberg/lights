// Command lights provides a command-line interface to control my two Elgato Key lights.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mdlayher/keylight"
)

type State struct {
	Info      bool
	Circadian bool
	Toggle    bool
	Key       Light
	Fill      Light
}

type Light struct {
	Addr        string
	Brightness  signedNumber
	Temperature signedNumber
}

func newState(now time.Time) *State {
	s := &State{}

	flag.BoolVar(&s.Info, "i", false, "display the current status of an Elgato Key Light without changing its state")
	flag.BoolVar(&s.Circadian, "c", false, "calculate and set the appropriate circadian lighting values")

	flag.StringVar(&s.Key.Addr, "ak", "http://keylight:9123", "the address of the Key Light's HTTP API")
	flag.Var(&s.Key.Brightness, "bk", "set Key Light brightness to an absolute (between 0 and 100) or relative (-N or +N) percentage")
	flag.Var(&s.Key.Temperature, "tk", "set Key Light temperature to an absolute (between 2900 and 7000) or relative (-N or +N) degrees")

	flag.StringVar(&s.Fill.Addr, "af", "http://filllight:9123", "the address of the Fill Light's HTTP API")
	flag.Var(&s.Fill.Brightness, "bf", "set Fill Light brightness to an absolute (between 0 and 100) or relative (-N or +N) percentage")
	flag.Var(&s.Fill.Temperature, "tf", "set Fill Light temperature to an absolute (between 2900 and 7000) or relative (-N or +N) degrees")

	flag.Parse()

	if s.Circadian {
		s.setCircadianValues(now)
	}

	// Only toggle the light if no modification flags are set.
	s.Toggle = !s.Key.Brightness.set && !s.Key.Temperature.set &&
		!s.Fill.Brightness.set && !s.Fill.Temperature.set

	return s
}

func main() {
	s := newState(time.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.SetFlags(0)

	if err := s.handleLight(ctx, s.Key); err != nil {
		log.Fatal(err)
	}

	if err := s.handleLight(ctx, s.Fill); err != nil {
		log.Fatal(err)
	}
}

func (s *State) setCircadianValues(now time.Time) {

}

func (s *State) handleLight(ctx context.Context, light Light) error {
	c, err := keylight.NewClient(light.Addr, nil)
	if err != nil {
		return fmt.Errorf("failed to create Key Light client: %w", err)
	}

	d, err := c.AccessoryInfo(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch accessory info: %w", err)
	}

	lights, err := c.Lights(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch key lights: %w", err)
	}

	if s.Info {
		// Log info and don't modify any settings.
		logInfo(d, lights)
		return nil
	}

	for _, l := range lights {
		if light.Brightness.relative {
			l.Brightness += light.Brightness.number
		} else if light.Brightness.set {
			l.Brightness = light.Brightness.number
		}
		if light.Temperature.relative {
			l.Temperature += light.Temperature.number
		} else if light.Temperature.set {
			l.Temperature = light.Temperature.number
		}

		if s.Toggle {
			l.On = !l.On
		} else {
			// If the light is being modified, force it on.
			l.On = true
		}
	}

	if err := c.SetLights(ctx, lights); err != nil {
		log.Fatalf("failed to set lights: %v", err)
	}

	logInfo(d, lights)

	return nil
}

type signedNumber struct {
	set      bool
	relative bool
	number   int
}

func (p signedNumber) String() string {
	if !p.set {
		return ""
	}

	if p.relative {
		return fmt.Sprintf("%+d", p.number)
	}

	return fmt.Sprintf("%d", p.number)
}

func (p *signedNumber) Set(s string) error {
	*p = signedNumber{}

	if s == "" {
		return nil
	}

	p.set = true
	negative := false

	if s[0] == '-' || s[0] == '+' {
		p.relative = true
		negative = s[0] == '-'
		s = s[1:]
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if negative {
		p.number = -n
	} else {
		p.number = n
	}

	return nil
}

// logInfo logs information about a device and its lights.
func logInfo(d *keylight.Device, ls []*keylight.Light) {
	name := d.DisplayName

	for _, l := range ls {
		onOff := "🚫"

		if l.On {
			onOff = "💡"
		}

		log.Printf("%s %s %dK %d%%", onOff, name, l.Temperature, l.Brightness)
	}
}
