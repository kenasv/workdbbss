//Основной пакет программы по логике её функционирования
package main

//Библиотеки, которые используются в этом пакете
import (
	"log"

	"github.com/go-pg/pg"

	//	"os"
	db "./db"
)

//Основная функция с которой стартует программа
func main() {

	log.Printf("Proga start.\n")
	pg_db := db.Connect()
	//SavePL(pg_db)
	db.PlaceHoldSelect(pg_db)
}

func SavePL(dbRef *pg.DB) {

	newPL1 := &db.Dostup101{
		ID:             "018",
		Diapazon:       "2600",
		Adres:          "Vladimir",
		Infrastructura: "Stolb",
		Prinadlegnost:  "MTS",
		Dostup:         "free",
		Kluch:          "Mottura",
		Aparatnaa:      "Intercross",
		AD:             "Vasily Pupkin",
		Contact:        "8-900-800-90-90",
		WGS:            "54.12345, 43.12345",
		Primechanie:    "test2",
		//IsActive:       true,
	}
	newPL2 := &db.Dostup101{
		ID:             "019",
		Diapazon:       "2600",
		Adres:          "Vladimir",
		Infrastructura: "Stolb",
		Prinadlegnost:  "MTS",
		Dostup:         "free",
		Kluch:          "Mottura",
		Aparatnaa:      "Intercross",
		AD:             "Vasily Pupkin",
		Contact:        "8-900-800-90-90",
		WGS:            "54.12345, 43.12345",
		Primechanie:    "test2",
		//IsActive:       true,
	}
	//newPL.SaveReturn(dbRef)
	totalItems := []*db.Dostup101{newPL1, newPL2}
	newPL1.SaveMulti(dbRef, totalItems)
}
