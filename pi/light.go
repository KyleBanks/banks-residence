package main

import (
	"github.com/nathan-osman/go-rpigpio"
)

const (
	gpioPinNumber = 4
)

type State rpi.Value

var (
	On  State = rpi.HIGH
	Off State = rpi.LOW
)

type Light struct {
	rpi.Pin

	State State
}

func newLight(pinNo int) (*Light, error) {
	p, err := rpi.OpenPin(pinNo, rpi.OUT)
	if err != nil {
		return nil, err
	}

	return &Light{Pin: *p, State: Off}, nil
}

func (l *Light) toggleState() {
	if l.State == Off {
		l.setState(On)
	} else {
		l.setState(Off)
	}
}

func (l *Light) setState(s State) {
	l.State = s
	l.Write(rpi.Value(s))
}
