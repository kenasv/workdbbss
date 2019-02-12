package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"workdbbss/db"

	"github.com/go-pg/pg"
)

func News(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/new.html")
		t.Execute(w, nil)
		fmt.Println("method: insert Pl")
	} else {
		r.ParseForm()
		fmt.Println("namePl:", r.Form["namePl"])

		Xaxaxa := r.FormValue("namePl")
		fmt.Println(Xaxaxa)
		pg_db := db.Connect()
		SavePL1(pg_db, Xaxaxa)
	}

}

func Del(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/new.html")
		t.Execute(w, nil)
		fmt.Println("method: delete Pl")
	} else {
		r.ParseForm()
		fmt.Println("namePl:", r.Form["namePl"])

		Xaxaxa := r.FormValue("namePl")
		fmt.Println(Xaxaxa)
		pg_db := db.Connect()
		DeletePl(pg_db, Xaxaxa)
	}

}

func DeletePl(dbRef *pg.DB, one string) {
	tvo := one
	newPL1 := &db.Dostup101{
		ID: tvo,
	}
	newPL1.DeletePl(dbRef)
}

func SavePL1(dbRef *pg.DB, one string) {

	tvo := one
	newPL1 := &db.Dostup101{
		ID:             tvo,
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
	newPL1.Save(dbRef)
}
