package main

import (
	"sync"

	"dfeprado.dev/rpg-master/api/master"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go master.RunMasterServer(wg)

	wg.Wait()
}
