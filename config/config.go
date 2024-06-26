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
	CentralServerPort string
	SleepTime         time.Duration
	UDPPort           string
}

// Environment variables
const (
	// EnvHTTPPort -
	EnvURL                    = "CP_URL"
	EnvHTTPPort               = "HTTP_PORT"
	EnvHTTPPortDefaultValue   = "8080"
	EnvGPIOListenList         = "GPIO_LISTEN"
	EnvGPIOSendList           = "GPIO_SEND"
	EnvFetchInterval          = "FETCH_INTERVAL"
	EnvCBIdentity             = "CB_IDENTITY"
	EnvCSPort                 = "CS_PORT"
	EnvCSPortPortDefaultValue = "12811"
	EnvSleepTime              = "LOCK_TIME"
	EnvSleepTimeDefaultValue  = "1h"
	EnvUDPPort                = "UDP_PORT"
	EnvUDPPortDefaultValue    = "1338"
)

// NewConfig will init env struct
func NewConfig() *Config {
	log.Printf("[INFO] Loading config for service...")
	config := viper.New()
	config.SetConfigName("service-config")
	config.AddConfigPath(".")
	config.AddConfigPath("/home/pi/")
	if err := config.ReadInConfig(); err != nil {
		log.Printf("[WARNING] [CONFIG] [SERVICE] %s", err.Error())
		log.Printf("[WARNING] [CONFIG] [SERVICE] Falling back to local env variables")
		config.AutomaticEnv()
	}
	config.SetDefault(EnvHTTPPort, EnvHTTPPortDefaultValue)
	config.SetDefault(EnvCSPort, EnvCSPortPortDefaultValue)
	config.SetDefault(EnvSleepTime, EnvSleepTimeDefaultValue)
	config.SetDefault(EnvUDPPort, EnvUDPPortDefaultValue)
	gpioListenRaw := config.GetString(EnvGPIOListenList)
	gpioSendRaw := config.GetString(EnvGPIOSendList)
	return &Config{
		URL:               config.GetString(EnvURL),
		ListenHTTPPort:    config.GetString(EnvHTTPPort),
		GPIOListenList:    strings.Split(gpioListenRaw, ","),
		GPIOSendList:      strings.Split(gpioSendRaw, ","),
		FetchInterval:     config.GetDuration(EnvFetchInterval),
		ChargeBoxIdentity: config.GetString(EnvCBIdentity),
		CentralServerPort: config.GetString(EnvCSPort),
		SleepTime:         config.GetDuration(EnvSleepTime),
		UDPPort:           config.GetString(EnvUDPPort),
	}
}
