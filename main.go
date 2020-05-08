package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Zarathos94/ocpp-service/config"
	"github.com/Zarathos94/ocpp-service/services/gpio"
	"github.com/Zarathos94/ocpp-service/services/listener"
	"github.com/Zarathos94/ocpp-service/services/ocpp"
	"github.com/Zarathos94/ocpp-service/services/point"
)

var (
	cfg = config.NewConfig()
	g   = gpio.NewIOService(cfg)
	cp  = point.NewCPointInterface(cfg)
	ls  = listener.NewListener(cfg)
)

// ListenToPins -
func ListenToPins(s *gpio.IOService) {
	go func() {
		for {
			select {
			case p := <-s.ListenChannel:
				go func() {
					log.Printf("[INFO] Got event on pin: %d", p)
					_, err := cp.ChangeAvailability(1, ocpp.AvailabilityType("Operative"))
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						return
					}
					// TODO: Add per pin specific actions
					_, err = cp.RemoteStartTransaction("someTag", 1)
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						return
					}
					for pin := range s.SendPins {
						s.SendSignalTimed(pin, s.Config.SleepTime)
					}
				}()
			}
		}
	}()
}

// ListenToUDP -
func ListenToUDP(s *listener.Listener) {
	go func() {
		for {
			select {
			case msg := <-s.ListenerChannel:
				a1 := strings.Split(msg, ",")
				action := a1[0]
				params := a1[1 : len(a1)-1]
				switch action {
				case "RESET":
					// RESET,[Hard|Soft]
					_, err := cp.Reset(ocpp.ResetType(params[0]))
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						return
					}
					for pin := range g.SendPins {
						g.SendSignalTimed(pin, 6*time.Minute)
					}
				case "AVAILABILITY":
					// AVAILABILITY,[connectorID,AvailabityType(Inoperative,Operative)]
					connID, _ := strconv.Atoi(params[0])
					_, err := cp.ChangeAvailability(int32(connID), ocpp.AvailabilityType(params[1]))
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						return
					}
				case "START":
					// START,[connectorID,tag(any string)]
					connID, _ := strconv.Atoi(params[0])
					_, err := cp.RemoteStartTransaction(params[1], int32(connID))
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						return
					}
				case "INVALIDATE":
					// INVALIDATE
					_, err := cp.Reset(ocpp.ResetType(params[0]))
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						return
					}
					for pin := range g.SendPins {
						g.SendSignalPersistent(pin)
					}
				}
			}
		}
	}()
}

func main() {

	//cp.GetRemoteConfiguration()
	if err := g.SetUp(); err != nil {
		log.Fatal(err)
	}
	if err := ls.Init(); err != nil {
		log.Fatal(err)
	}
	go ls.Start()
	g.Start()
	defer g.Stop()
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM)
	signal.Notify(syscallCh, syscall.SIGINT)
	signal.Notify(syscallCh, syscall.SIGKILL)
	<-syscallCh
}
