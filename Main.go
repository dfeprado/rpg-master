package main

import (
	"fmt"
	"os"
	"slices"
	"sync"

	"dfeprado.dev/rpg-master/api/master"
)

func main() {
	fmt.Print("RPG Master")
	if slices.Contains(os.Args, "--dev") {
		fmt.Println(" DEV MODE")
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go master.RunMasterServer(wg)

	wg.Wait()
}
