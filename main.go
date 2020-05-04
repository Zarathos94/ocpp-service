package main

import (
	"log"
	"time"

	"github.com/Zarathos94/ocpp-service/config"
	"github.com/Zarathos94/ocpp-service/services/ocpp"
	"github.com/Zarathos94/ocpp-service/services/point"
)

func main() {
	cp := point.NewCPointInterface(config.NewConfig())
	resp, err := cp.GetRemoteConfiguration()
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	for k, v := range resp {
		log.Printf("%s=%s", k, v)
	}
	_, err = cp.ChangeRemoteConfiguration("LOCALAUTHORIZATIONLISTENABLED", "false")
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	resp1, err := cp.GetRemoteConfiguration()
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	for k, v := range resp1 {
		log.Printf("%s=%s", k, v)
	}
	_, err = cp.ChangeAvailability(1, ocpp.AvailabilityType("Operative"))
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.ChangeRemoteConfiguration("Jack", "H")
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.ClearCache()
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.Reset(ocpp.ResetType("Hard"))
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.RemoteStartTransaction("someTag", 1)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.RemoteStopTransaction(1)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.ReserveNow(1, time.Now().Add(10*time.Minute), "someTag", "someParentTag", 1)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.CancelReservation(1)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.UnlockConnector(1)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	_, err = cp.DataTransfer("someVendor", "1", "some message")
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	/*headers := struct {
		ChargeBoxIdentity string
		Action            string
	}{
		ChargeBoxIdentity: "boxid",
		Action:            "/GetConfiguration",
	}

	cli := soap.Client{
		URL:                    uri,
		Namespace:              ocpp.Namespace,
		URNamespace:            ocpp.Namespace,
		ExcludeActionNamespace: false,
		ContentType:            "application/soap+xml;charset=UTF-8",
		Header:                 headers,
	}
	soapService := ocpp.NewChargePointService(&cli)
	conf, err := soapService.GetConfiguration(&ocpp.GetConfigurationRequest{})
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	for _, v := range conf.ConfigurationKey {
		log.Printf("%v", *v)
	}
	headers.Action = "/RemoteStartTransaction"
	cli.Header = headers
	resp1, err := soapService.RemoteStartTransaction(&ocpp.RemoteStartTransactionRequest{
		IdTag:       ocpp.IdToken("Simulator 1"),
		ConnectorId: 1,
	})
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	log.Printf("%v", resp1)*/
}
