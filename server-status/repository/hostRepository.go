package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"../models"
)

/*
get database connection
*/
func connection() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", config.getUrlConnection("postgres"))
	if err != nil {
		if db != nil {
			_ = db.Close()
		}
	}
	return
}

func Save(host *models.Host) (int64, error) {
	db, err := connection()
	if err != nil {
		log.Fatal(err)
	}

	entity, err := db.Exec(`INSERT INTO host (hostname) VALUES ($1);`, host.Hostname)
	if err != nil {
		log.Fatal(err)
	}
	return entity.LastInsertId()
}
