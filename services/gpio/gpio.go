package gpio

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/Zarathos94/ocpp-service/services/ocpp"
	"github.com/Zarathos94/ocpp-service/services/point"

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
	PointService  *point.CPointInterface
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
	//go io.pinWatcher2(pin, &io.ListenChannel)
	go func(p int64) {
		log.Printf("[INFO] Listening to events on pin: %d", pin)
		//io.SendPins[pin].Detect(rpio.AnyEdge)

		for {
			if io.ListenPins[p].EdgeDetected() {
				log.Printf("Edge detected: %d", p)
				//io.ListenPins[p].High()
				//io.ListenPins[p].Detect(rpio.NoEdge)
				//io.ListenPins[p].Input()
				io.ListenPins[p].PullUp()
				io.ListenPins[p].Detect(rpio.FallEdge)
				io.ListenChannel <- p
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(pin)
}

// Plan B for a while...
func (io *IOService) pinWatcher2(pinNum int64, ch *chan int64) {
	pin := rpio.Pin(uint8(pinNum))
	pin.Input()
	pin.PullUp()
	//lastTime := time.Now().UnixNano() / 1000000
	lastState := rpio.High
	for {
		state := pin.Read()
		if state == rpio.High {
		} else if state == rpio.Low {
			if lastState == rpio.High {
				pin.Low()
				//now := time.Now().UnixNano() / 1000000
				//diff := now - lastTime
				//lastTime = now
				*ch <- pinNum
			}
		}
		lastState = state
		time.Sleep(time.Millisecond)
	}
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
		io.SendPins[pin].High()

		io.MapLock.Lock()
		io.LockedPins[pin] = false
		delete(io.LockedPins, pin)
		io.MapLock.Unlock()
		time.Sleep(duration)
		io.SendPins[pin].Low()
		_, err := io.PointService.ChangeAvailability(1, ocpp.AvailabilityTypeInoperative)
		if err != nil {
			log.Printf("[ERROR] Could not change availability")
			log.Printf("SoapClient error: %s", err)
			return
		}
	}()
	return true
}

// SendSignalPersistent -
func (io *IOService) SendSignalPersistent(pin int64) bool {
	io.SendPins[pin].Toggle()
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
		//io.ListenPins[i].High()
		io.ListenPins[i].PullUp()
		io.ListenPins[i].Detect(rpio.FallEdge)
	}
	for _, v := range io.Config.GPIOSendList {
		i, _ := strconv.ParseInt(v, 10, 64)
		io.SendPins[i] = rpio.Pin(i)
		io.SendPins[i].Output()
		io.SendPins[i].Low()
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
