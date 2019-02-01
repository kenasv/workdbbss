package db

import (
	"log"

	"github.com/go-pg/pg"
)

func PlaceHoldSelect(db *pg.DB) error {

	var value int
	var query string = "SELECT ?1"
	_, selectErr := db.Query(pg.Scan(&value), query, 42, 41)
	if selectErr != nil {

		log.Printf("Eror while Select running query, Reason: %v\n", selectErr)
		return selectErr
	}

	log.Printf("Scan value complite, scsn value: %d\n", value)
	return nil

}
