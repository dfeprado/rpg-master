package repository

import (
	"dfeprado.dev/rpg-master/api/common/database"
	"dfeprado.dev/rpg-master/api/common/model"
)

func SelectAllWeapons() ([]*model.Weapon, error) {
	rows, err := database.OpenDB().Query(
		`select
			id,
			name,
			weight,
			pa,
			damage,
			two_handed,
			distance,
			effect,
			size,
			enhancement
		from weapons;
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	weapons := make([]*model.Weapon, 0)
	for rows.Next() {
		var weapon model.Weapon
		if err := rows.Scan(
			&weapon.Item.Id,
			&weapon.Item.Name,
			&weapon.Item.Weight,
			&weapon.PA,
			&weapon.Damage,
			&weapon.TwoHanded,
			&weapon.Distance,
			&weapon.Effect,
			&weapon.Size,
			&weapon.Enhancement,
		); err != nil {
			return nil, err
		}
		weapons = append(weapons, &weapon)
	}

	return weapons, nil
}

func SaveWeapon(weapon *model.Weapon) error {
	db := database.OpenDB()

	if weapon.Item.Id != 0 {
		_, err := db.Exec(
			`update weapons set
				name = ?,
				weight = ?,
				pa = ?,
				damage = ?,
				two_handed = ?,
				distance = ?
				effect = ?,
				size = ?,
				enhancement = ?
			where
				id = ?`,
			weapon.Item.Name,
			weapon.Item.Weight,
			weapon.PA,
			weapon.Damage,
			weapon.TwoHanded,
			weapon.Distance,
			weapon.Effect,
			weapon.Enhancement,
			weapon.Item.Id,
		)

		return err
	} else {
		nextId, err := database.GetNextId("weapons")
		if err != nil {
			return err
		}

		_, err = db.Exec(
			`insert into weapons(
				id, name, weight, pa, damage, two_handed, distance, effect, size, enhancement
			) 
			values (
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?
			)`,
			nextId,
			weapon.Item.Name,
			weapon.Item.Weight,
			weapon.PA,
			weapon.Damage,
			weapon.TwoHanded,
			weapon.Distance,
			weapon.Effect,
			weapon.Enhancement,
			weapon.Item.Id,
		)

		if err == nil {
			weapon.Item.Id = nextId
		}

		return err
	}
}
