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
		// http.SetCookie(w, &http.Cookie{Name: "UI", Value: "MASTER"})
		fmt.Fprintln(w, "{\"hello\": \"world\"}")
	})
	apiRoutes := api.NewHandler(routes)

	// TODO discover the next available port
	// TODO set the same port as players server
	address := "127.0.0.1:8080"
	fmt.Printf("You (master) can connect through %s\n", address)
	http.ListenAndServe(address, apiRoutes)
	wg.Done()
}
