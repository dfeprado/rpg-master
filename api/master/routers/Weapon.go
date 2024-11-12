package routers

import (
	"dfeprado.dev/rpg-master/api"
	"dfeprado.dev/rpg-master/api/common/repository"
)

func GetSelectAllWeapons(ctx *api.Context) any {
	weapons, err := repository.SelectAllWeapons()
	if err != nil {
		// TODO tratar
	}
	return weapons
}
