package master

import (
	"fmt"

	"dfeprado.dev/rpg-master/api"
)

func SetRouters(routers *api.Router) {
	routers.Get("/", func(ctx *api.Context) any {
		type Person struct {
			Name string `json:"name"`
			Age  uint8  `json:"age"`
		}

		response := &Person{"Foo", 32}
		return response
	})
	routers.Get("/none", func(ctx *api.Context) any {
		fmt.Printf("Ok\n")
		return nil
	})
	routers.Get("/err", func(ctx *api.Context) any {
		panic("Err path")
	})
}
