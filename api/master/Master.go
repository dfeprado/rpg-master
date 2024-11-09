package master

import (
	"fmt"
	"net/http"
	"sync"

	"dfeprado.dev/rpg-master/api"
)

func RunMasterServer(wg *sync.WaitGroup) {
	routes := http.NewServeMux()
	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/json")
		fmt.Fprintln(w, "{\"hello\": \"world\"}")
	})
	app := api.GetApplication()
	apiRoutes := api.NewHandler(routes, app)

	// TODO discover the next available port
	address := fmt.Sprintf("127.0.0.1:%d", app.GetPort())
	fmt.Printf("You (master) can connect through http://%s\n", address)
	http.ListenAndServe(address, apiRoutes)
	wg.Done()
}
