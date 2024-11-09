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
	router.Get("/", func(ctx *api.Context) any {
		type Person struct {
			Name string `json:"name"`
			Age  uint8  `json:"age"`
		}

		response := &Person{"Foo", 32}
		return response
	})
	router.Get("/none", func(ctx *api.Context) any {
		fmt.Printf("Ok\n")
		return nil
	})

	// TODO discover the next available port
	address := fmt.Sprintf("127.0.0.1:%d", app.GetPort())
	fmt.Printf("You (master) can connect through http://%s\n", address)
	http.ListenAndServe(address, router)
	wg.Done()
}
