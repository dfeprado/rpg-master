package main

import (
	"fmt"
	"os"
	"slices"
	"sync"

	"dfeprado.dev/rpg-master/api/master"
	"dfeprado.dev/rpg-master/api/player"
)

func main() {
	fmt.Print("RPG Master")
	if slices.Contains(os.Args, "--dev") {
		fmt.Println(" DEV MODE")
	}
	fmt.Println()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go master.RunMasterServer(wg)
	go player.RunPlayerServer(wg)

	wg.Wait()
}
