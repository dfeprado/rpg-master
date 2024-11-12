package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func OpenDB() *sql.DB {
	if db == nil {
		_db, err := sql.Open("sqlite", "./db/game.sdb")
		if err != nil {
			panic(fmt.Sprintf("database: Error while opening: %v\n", err))
		}

		db = _db
	}

	return db
}

func GetNextId(table string) (uint64, error) {
	row := db.QueryRow("select max(id) from " + table)
	var lastId uint64
	if err := row.Scan(&lastId); err != nil {
		if !strings.Contains(err.Error(), "converting NULL to uint64") {
			return 0, err
		}
	}

	return lastId + 1, nil
}
