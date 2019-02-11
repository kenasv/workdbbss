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
		fmt.Println("method: 1")
	} else {
		r.ParseForm()
		fmt.Println("namePl:", r.Form["namePl"])

		
		Xaxaxa := r.FormValue("namePl")
		fmt.Println(Xaxaxa)
		pg_db := db.Connect()
		DeletePl(pg_db)
	}

}

func DeletePl(dbRef *pg.DB) {
	newPL1 := &db.Dostup101{
		Diapazon: "1800",
	}
	newPL1.DeletePl(dbRef)
}
