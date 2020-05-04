package point

import (
	"github.com/Zarathos94/ocpp-service/config"
	"github.com/Zarathos94/ocpp-service/services/ocpp"
	"github.com/hooklift/gowsdl/soap"
)

type internalHeaders struct {
	ChargeBoxIdentity string
	Action            string
}

// CPointInterface -
type CPointInterface struct {
	Config     *config.Config
	Headers    map[string]string
	OCPPClient *soap.Client
	CPService  ocpp.ChargePointService
	intHeaders internalHeaders
}

// NewCPointInterface -
func NewCPointInterface(cfg *config.Config) *CPointInterface {
	headers := map[string]string{
		"Content-Type": "application/soap+xml;charset=UTF-8",
	}
	cli := soap.NewClient(cfg.URL, soap.WithHTTPHeaders(headers))
	return &CPointInterface{
		Headers:    headers,
		OCPPClient: cli,
		CPService:  ocpp.NewChargePointService(cli),
		//intHeaders: headers,
	}
}

func (cp *CPointInterface) preSetActionHeaders(cmd string) {
	cp.intHeaders.Action = cmd
	//cp.OCPPClient.Header = cp.intHeaders
}
