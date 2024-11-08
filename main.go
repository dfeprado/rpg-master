package main

import (
	"fmt"
	"sync"

	"dfeprado.dev/rpg-master/rpgmaster/master"
)

func main() {
	fmt.Println("RPG-Master ALPHA")
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// go net.StartPlayerServer(wg)
	go master.StartServer(wg)

	wg.Wait()
}
