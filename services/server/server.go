package server

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Zarathos94/ocpp-service/config"
	oc "github.com/eduhenke/go-ocpp"
	ocs "github.com/eduhenke/go-ocpp/cs"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
)

// CentralServer -
type CentralServer struct {
	Config         *config.Config
	CentralService ocs.CentralSystem
}

// NewCentralServer -
func NewCentralServer(cfg *config.Config) *CentralServer {
	return &CentralServer{
		Config:         cfg,
		CentralService: ocs.New(),
	}
}

// StartListening -
func (c *CentralServer) StartListening() {
	oc.SetDebugLogger(log.New(os.Stdout, "DEBUG:", log.Ltime))
	oc.SetErrorLogger(log.New(os.Stderr, "ERROR:", log.Ltime))
	go c.CentralService.Run(":"+c.Config.CentralServerPort, func(req cpreq.ChargePointRequest, cpID string) (cpresp.ChargePointResponse, error) {
		fmt.Printf("EXAMPLE(MAIN): Request from %s\n", cpID)
		switch req := req.(type) {
		case *cpreq.BootNotification:
			return &cpresp.BootNotification{
				Status:      "Accepted",
				CurrentTime: time.Now(),
				Interval:    60,
			}, nil
		case *cpreq.Authorize:
			expire := time.Now().Add(10 * time.Minute)
			return &cpresp.Authorize{
				IdTagInfo: &cpresp.IdTagInfo{
					ExpiryDate:  &expire,
					ParentIdTag: "",
					Status:      "Accepted",
				},
			}, nil
		case *cpreq.Heartbeat:
			fmt.Printf("EXAMPLE(MAIN): Heartbeat\n")
			return &cpresp.Heartbeat{CurrentTime: time.Now()}, nil

		case *cpreq.StatusNotification:
			if req.Status != "Available" {
				// chargepoint is unavailable
			}
			return &cpresp.StatusNotification{}, nil

		default:
			fmt.Printf("EXAMPLE(MAIN): action not supported: %s\n", req.Action())
			return nil, errors.New("Response not supported")
		}

	})
}
