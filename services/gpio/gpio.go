package gpio

import (
	"errors"

	"github.com/Zarathos94/ocpp-service/config"
	"github.com/stianeikeland/go-rpio/v4"
)

// IOMode -
type IOMode uint8

// IOMode definitions
const (
	Input IOMode = iota
	Output
	Clock
	Pwm
)

// IOService -
type IOService struct {
	Config *config.Config
	Pins   map[int]rpio.Pin
}

// NewIOService -
func NewIOService(cfg *config.Config) *IOService {
	return &IOService{
		Config: cfg,
		Pins:   make(map[int]rpio.Pin),
	}
}

// Start -
func (io *IOService) Start() error {
	err := rpio.Open()
	return err
}

// SetMode -
func (io *IOService) SetMode(pin int, mode IOMode) error {
	if _, ok := io.Pins[pin]; !ok {
		io.Pins[pin] = rpio.Pin(pin)
	}
	switch mode {
	case Input:
		io.Pins[pin].Input()
	case Output:
		io.Pins[pin].Output()
	case Clock:
		io.Pins[pin].Clock()
	case Pwm:
		io.Pins[pin].Pwm()
	default:
		return errors.New("Wrong mode")
	}
	return nil
}

// Stop -
func (io *IOService) Stop() error {
	return rpio.Close()
}
