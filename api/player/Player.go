package player

import (
	"fmt"
	"net/http"
	"sync"

	"dfeprado.dev/rpg-master/api"
)

func RunPlayerServer(wg *sync.WaitGroup) {
	routes := http.NewServeMux()
	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"Hello\": \"Player\"}")
	})
	app := api.GetApplication()
	apiRoutes := api.NewHandler(routes, app)

	fmt.Printf("Your players can connect through http://%s\n", app.JoinHostAndPort())
	http.ListenAndServe(app.JoinHostAndPort(), apiRoutes)

	wg.Done()
}
