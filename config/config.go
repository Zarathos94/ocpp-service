package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config -
type Config struct {
	URL               string
	ChargeBoxIdentity string
	ListenHTTPPort    string
	GPIOListenList    []string
	GPIOSendList      []string
	FetchInterval     time.Duration
}

// Environment variables
const (
	// EnvHTTPPort -
	EnvURL                  = "CP_URL"
	EnvHTTPPort             = "HTTP_PORT"
	EnvHTTPPortDefaultValue = "8080"
	EnvGPIOListenList       = "GPIO_LISTEN"
	EnvGPIOSendList         = "GPIO_SEND"
	EnvFetchInterval        = "FETCH_INTERVAL"
	EnvCBIdentity           = "CB_IDENTITY"
)

// NewConfig will init env struct
func NewConfig() *Config {
	log.Printf("[INFO] Loading config for service...")
	config := viper.New()
	config.SetConfigName("service-config")
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		log.Printf("[WARNING] [CONFIG] [SERVICE] %s", err.Error())
		log.Printf("[WARNING] [CONFIG] [SERVICE] Falling back to local env variables")
		config.AutomaticEnv()
	}
	config.SetDefault(EnvHTTPPort, EnvHTTPPortDefaultValue)

	gpioListenRaw := config.GetString(EnvGPIOListenList)
	gpioSendRaw := config.GetString(EnvGPIOSendList)
	return &Config{
		URL:               config.GetString(EnvURL),
		ListenHTTPPort:    config.GetString(EnvHTTPPort),
		GPIOListenList:    strings.Split(gpioListenRaw, ","),
		GPIOSendList:      strings.Split(gpioSendRaw, ","),
		FetchInterval:     config.GetDuration(EnvFetchInterval),
		ChargeBoxIdentity: config.GetString(EnvCBIdentity),
	}
}
