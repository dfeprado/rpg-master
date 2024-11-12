package master

import (
	"dfeprado.dev/rpg-master/api"
	"dfeprado.dev/rpg-master/api/master/routers"
)

func SetRouters(masterRouters *api.Router) {
	masterRouters.Get("/", func(ctx *api.Context) any {
		type Person struct {
			Name string `json:"name"`
			Age  uint8  `json:"age"`
		}

		response := &Person{"Foo", 32}
		return response
	})

	masterRouters.Get("/weapons", routers.GetSelectAllWeapons)
}
