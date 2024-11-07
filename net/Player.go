package net

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

type PlayerConnection struct {
	Ip string
	Port string
}

func (p *PlayerConnection) GetHostAndPort() string {
	return fmt.Sprintf("%s:%s", p.Ip, p.Port)
}

var playerConnection *PlayerConnection = nil

func GetPlayerConnectionStruct() *PlayerConnection {
	if playerConnection == nil {
		conn, err := net.Dial("udp", "1.1.1.1:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		ip, _, err := net.SplitHostPort(conn.LocalAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		playerConnection = &PlayerConnection{
			ip,
			"8080",
		}
	}

	return playerConnection
}

func StartPlayerServer(wg *sync.WaitGroup) {
	pconn := GetPlayerConnectionStruct()
	address := pconn.GetHostAndPort()
	fmt.Printf(
		"Your players can connect through http://%s address\n", 
		address,
	)
	log.Fatal(http.ListenAndServe(address, &_HTTPHandler{
		func (w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, my fellow player!")
		},
	}))
	wg.Done()
}