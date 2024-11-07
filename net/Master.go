package net

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func StartMasterServer(wg *sync.WaitGroup) {
	address := "127.0.0.1:8080"
	fmt.Printf("You (master) can connect through http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, &_HTTPHandler{
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "<p>Hello, master!</p>")
			playerAddress := GetPlayerConnectionStruct().GetHostAndPort()
			fmt.Fprintf(w, "<p>Your players can connect through <a href=\"http://%s\" target=\"_blank\">http://%s</a> address<p>", playerAddress, playerAddress)
		},
	}))
	wg.Done()
}
