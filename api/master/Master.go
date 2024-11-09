package master

import (
	"fmt"
	"net/http"
	"sync"

	"dfeprado.dev/rpg-master/api"
)

func RunMasterServer(wg *sync.WaitGroup) {
	app := api.GetApplication()
	router := api.NewRouter(app)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "index of api")
	})

	// TODO discover the next available port
	address := fmt.Sprintf("127.0.0.1:%d", app.GetPort())
	fmt.Printf("You (master) can connect through http://%s\n", address)
	http.ListenAndServe(address, router)
	wg.Done()
}
