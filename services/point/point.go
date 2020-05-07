package point

import (
	"github.com/Zarathos94/ocpp-service/config"
)

type internalHeaders struct {
	ChargeBoxIdentity string
	Action            string
}

// CPointInterface -
type CPointInterface struct {
	Config *config.Config
}

// NewCPointInterface -
func NewCPointInterface(cfg *config.Config) *CPointInterface {
	return &CPointInterface{
		Config: cfg,
	}
}
