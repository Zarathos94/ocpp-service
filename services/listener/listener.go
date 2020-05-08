package listener

import (
	"log"
	"net"
	"strings"

	"github.com/Zarathos94/ocpp-service/config"
)

// Listener -
type Listener struct {
	Config          *config.Config
	addr            *net.UDPAddr
	udpConn         *net.UDPConn
	ListenerChannel chan string
}

// NewListener -
func NewListener(cfg *config.Config) *Listener {
	return &Listener{
		Config:          cfg,
		ListenerChannel: make(chan string),
	}
}

// Init -
func (l *Listener) Init() error {
	s, err := net.ResolveUDPAddr("udp4", ":"+l.Config.UDPPort)
	if err != nil {
		return err
	}
	l.addr = s
	connection, err := net.ListenUDP("udp4", l.addr)
	if err != nil {
		return err
	}
	l.udpConn = connection
	return err
}

// Start -
func (l *Listener) Start() {

	defer l.udpConn.Close()
	buffer := make([]byte, 1024)
	log.Printf("Listening on UDP Port: %s", l.Config.UDPPort)
	for {
		n, addr, err := l.udpConn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}
		if len(buffer) == 0 {
			continue
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" || strings.TrimSpace(string(buffer[0:n])) == "RESTART" {
			log.Printf("Exiting UDP server!")
			return
		}
		l.ListenerChannel <- string(buffer[0 : n-1])

		log.Printf("data: OK\n")
		_, err = l.udpConn.WriteToUDP([]byte("OK"), addr)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
	}
}
