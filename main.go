package main

import (
	"log"

	"github.com/Zarathos94/ocpp-service/config"
	"github.com/Zarathos94/ocpp-service/services/gpio"
)

func main() {

	/*cp := point.NewCPointInterface(config.NewConfig())

	resp, err := cp.GetRemoteConfiguration()
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	for k, v := range resp {
		log.Printf("%s=%s", k, v)
	}
	_, err = cp.ChangeAvailability(1, ocpp.AvailabilityType("Operative"))
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.RemoteStartTransaction("someTag", 1)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}*/
	g := gpio.NewIOService(config.NewConfig())
	if err := g.Start(); err != nil {
		log.Fatal(err)
	}
	g.SetMode(16, gpio.Output)
	g.Pins[16].High()
}
