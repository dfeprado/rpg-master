package main

import (
	"fmt"
	"sync"

	"dfeprado.dev/rpg-master/net"
)

func main() {
	fmt.Println("RPG-Master ALPHA")
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go net.StartPlayerServer(wg)
	go net.StartMasterServer(wg)

	wg.Wait()
}
