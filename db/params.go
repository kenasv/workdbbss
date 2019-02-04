package db

import (
	"log"

	"github.com/go-pg/pg"
)

type Params struct {
	Param1 string
	Param2 string
}

func PlaceHoldSelect(db *pg.DB) error {

	var value string
	params := Params{
		Param1: "This param 1",
		Param2: "This param 2",
	}

	var query string = "SELECT ?2"
	_, selectErr := db.Query(pg.Scan(&value), query, params)
	if selectErr != nil {

		log.Printf("Eror while Select running query, Reason: %v\n", selectErr)
		return selectErr
	}

	log.Printf("Scan value complite, scsn value: %d\n", value)
	return nil

}
