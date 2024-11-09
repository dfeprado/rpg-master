package api

import (
	"fmt"
	"log"
	"net"
	"os"
	"slices"
)

type Application struct {
	ip    string
	port  uint16
	isDev bool
}

func (p *Application) JoinHostAndPort() string {
	return fmt.Sprintf("%s:%d", p.ip, p.port)
}

func (p *Application) GetIp() string {
	return p.ip
}

func (p *Application) GetPort() uint16 {
	return p.port
}

func (p *Application) IsDev() bool {
	return p.isDev
}

var application *Application = nil

func GetApplication() *Application {
	if application == nil {
		conn, err := net.Dial("udp", "1.1.1.1:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		ip, _, err := net.SplitHostPort(conn.LocalAddr().String())
		if err != nil {
			log.Fatal(err)
		}

		// TODO do a port discovery, so that the application can
		// run even when the port is being used by another application
		application = &Application{
			ip,
			8080,
			slices.Contains(os.Args, "--dev"),
		}
	}

	return application
}
