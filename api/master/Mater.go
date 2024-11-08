package master

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"sync"
)

type devRoutes struct {
	Handler http.Handler
}

func (d *devRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
	d.Handler.ServeHTTP(w, r)
}

func RunMasterServer(wg *sync.WaitGroup) {
	routes := http.NewServeMux()
	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/json")
		fmt.Fprintln(w, "{\"hello\": \"world\"}")
	})
	apiRoutes := http.StripPrefix("/api", routes)
	if slices.Contains(os.Args, "--dev") {
		apiRoutes = &devRoutes{apiRoutes}
	}

	// TODO discover the next available port
	// TODO set the same port as players server
	address := "127.0.0.1:8080"
	fmt.Printf("You (master) can connect through %s\n", address)
	http.ListenAndServe(address, apiRoutes)
	wg.Done()
}
