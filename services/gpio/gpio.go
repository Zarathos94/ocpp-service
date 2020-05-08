package gpio

import (
	"log"
	"strconv"
	"sync"
	"time"

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
	Config        *config.Config
	ListenPins    map[int64]rpio.Pin
	SendPins      map[int64]rpio.Pin
	MapLock       *sync.Mutex
	LockedPins    map[int64]bool
	ListenChannel chan int64
}

// NewIOService -
func NewIOService(cfg *config.Config) *IOService {
	return &IOService{
		Config:        cfg,
		ListenPins:    make(map[int64]rpio.Pin),
		SendPins:      make(map[int64]rpio.Pin),
		ListenChannel: make(chan int64),
		LockedPins:    make(map[int64]bool),
		MapLock:       &sync.Mutex{},
	}
}

func (io *IOService) listenToPin(pin int64) {
	go func() {
		log.Printf("[INFO] Listening to events on pin: %d", pin)
		for {
			if io.SendPins[pin].EdgeDetected() {
				io.ListenChannel <- pin
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// SendSignalTimed -
func (io *IOService) SendSignalTimed(pin int64, duration time.Duration) bool {

	if val, ok := io.LockedPins[pin]; val && ok {
		return false
	}
	io.MapLock.Lock()
	io.LockedPins[pin] = true
	io.MapLock.Unlock()
	go func() {
		io.ListenPins[pin].High()

		io.MapLock.Lock()
		io.LockedPins[pin] = false
		delete(io.LockedPins, pin)
		io.MapLock.Unlock()
		time.Sleep(duration)
		io.ListenPins[pin].Low()
	}()
	return true
}

// SendSignalPersistent -
func (io *IOService) SendSignalPersistent(pin int64) bool {
	io.ListenPins[pin].Toggle()
	return true
}

// SetUp -
func (io *IOService) SetUp() error {
	err := rpio.Open()
	io.preparePins()
	return err
}

func (io *IOService) preparePins() {

	for _, v := range io.Config.GPIOListenList {
		i, _ := strconv.ParseInt(v, 10, 64)
		io.ListenPins[i] = rpio.Pin(i)
		io.ListenPins[i].Input()
		io.ListenPins[i].PullUp()
		io.ListenPins[i].Detect(rpio.FallEdge)
	}
	for _, v := range io.Config.GPIOSendList {
		i, _ := strconv.ParseInt(v, 10, 64)
		io.SendPins[i] = rpio.Pin(i)
		io.SendPins[i].Output()
	}
}

// Start -
func (io *IOService) Start() error {

	for i := range io.ListenPins {
		io.listenToPin(i)
	}
	return nil
}

// Stop -
func (io *IOService) Stop() error {
	return rpio.Close()
}
