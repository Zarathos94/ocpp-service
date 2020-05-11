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
	"github.com/stianeikeland/go-rpio"
)

var (
	cfg = config.NewConfig()
	g   = gpio.NewIOService(cfg)
	cp  = point.NewCPointInterface(cfg)
	ls  = listener.NewListener(cfg)
)

// ListenToPins -
func ListenToPins(s *gpio.IOService) {
	log.Printf("[INFO] Listening to pin events")
	go func() {
		for {
			select {
			case p := <-s.ListenChannel:
				if p == 0 {
					continue
				}
				go func(pin int64) {
					log.Printf("[INFO] Got event on pin: %d", pin)
					_, err := cp.ChangeAvailability(1, ocpp.AvailabilityTypeOperative)
					if err != nil {
						log.Printf("[ERROR] Could not change availability")
						log.Printf("SoapClient error: %s", err)
						return
					}
					time.Sleep(1 * time.Second)
					// TODO: Add per pin specific actions
					_, err = cp.RemoteStartTransaction("33334", 1)
					if err != nil {
						log.Printf("[ERROR] Could not start transaction")
						log.Printf("SoapClient error: %s", err)
						return
					}
					for pin := range s.SendPins {
						s.SendSignalTimed(pin, s.Config.SleepTime)
					}
				}(p)
			default:
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
				if msg == "" {
					continue
				}
				if len(msg) < 3 {
					continue
				}
				a1 := strings.Split(msg, ",")
				action := a1[0]
				params := a1[1 : len(a1)-1]
				switch action {
				case "RESET":
					// RESET,[Hard|Soft]
					if strings.ToLower(params[0]) == "hard" {
						_, err := cp.Reset(ocpp.ResetTypeHard)
						if err != nil {
							log.Printf("SoapClient error: %s", err)
							continue
						}
					}
					if strings.ToLower(params[0]) == "soft" {
						_, err := cp.Reset(ocpp.ResetTypeSoft)
						if err != nil {
							log.Printf("SoapClient error: %s", err)
							continue
						}
					}

					for pin := range g.SendPins {
						g.SendSignalTimed(pin, 5*time.Minute)
					}
					time.Sleep(6 * time.Minute)
					os.Exit(501)
				case "AVAILABILITY":
					// AVAILABILITY,[connectorID,AvailabityType(Inoperative,Operative)]
					connID, _ := strconv.Atoi(params[0])
					if strings.ToLower(params[1]) == "inoperative" {
						_, err := cp.ChangeAvailability(int32(connID), ocpp.AvailabilityTypeInoperative)
						if err != nil {
							log.Printf("SoapClient error: %s", err)
							continue
						}
					}
					if strings.ToLower(params[1]) == "operative" {
						_, err := cp.ChangeAvailability(int32(connID), ocpp.AvailabilityTypeOperative)
						if err != nil {
							log.Printf("SoapClient error: %s", err)
							continue
						}
					}
				case "START":
					// START,[connectorID,tag(any string)]
					connID, _ := strconv.Atoi(params[0])
					_, err := cp.RemoteStartTransaction(params[1], int32(connID))
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						continue
					}
				case "INVALIDATE":
					// INVALIDATE
					_, err := cp.Reset(ocpp.ResetTypeSoft)
					if err != nil {
						log.Printf("SoapClient error: %s", err)
						continue
					}
					for pin := range g.SendPins {
						g.SendSignalPersistent(pin)
					}
				}
			default:
			}
		}
	}()
}

func main() {

	if _, err := cp.GetRemoteConfiguration(); err != nil {
		log.Fatalf("Startup error: %v", err.Error())
	}
	if _, err := cp.ChangeAvailability(1, ocpp.AvailabilityTypeInoperative); err != nil {
		log.Fatalf("Startup error: %v", err.Error())
	}
	if _, err := cp.ChangeAvailability(2, ocpp.AvailabilityTypeInoperative); err != nil {
		log.Fatalf("Startup error: %v", err.Error())
	}
	if err := g.SetUp(); err != nil {
		log.Fatal(err)
	}
	if err := ls.Init(); err != nil {
		log.Fatal(err)
	}
	go ls.Start()
	g.Start()

	ListenToPins(g)
	ListenToUDP(ls)
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM)
	signal.Notify(syscallCh, syscall.SIGINT)
	signal.Notify(syscallCh, syscall.SIGKILL)
	<-syscallCh
	log.Printf("[INFO] Clearing state... Stopping service")
	for p := range g.ListenPins {
		g.ListenPins[p].Detect(rpio.NoEdge)
	}
	g.Stop()
}
