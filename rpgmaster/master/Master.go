package master

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"dfeprado.dev/rpg-master/rpgmaster"
	"dfeprado.dev/rpg-master/rpgmaster/master/ui"
)

func StartServer(wg *sync.WaitGroup) {
	app := rpgmaster.GetApplication()
	address := fmt.Sprintf("127.0.0.1:%d", app.GetPort())
	fmt.Printf("You (master) can connect through http://%s\n", address)

	var MasterMux *http.ServeMux = http.NewServeMux()

	MasterMux.HandleFunc("/", ui.Render)
	MasterMux.Handle("/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("rpgmaster/master/ui")),
		),
	)

	log.Fatal(http.ListenAndServe(address, MasterMux))
	wg.Done()
}
