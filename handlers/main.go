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
		one1 := r.FormValue("insert_ID")
		one2 := r.FormValue("insert_Diapazon")
		one3 := r.FormValue("insert_Adres")
		one4 := r.FormValue("insert_Infrastructura")
		one5 := r.FormValue("insert_Prinadlegnost")
		one6 := r.FormValue("insert_Dostup")
		one7 := r.FormValue("insert_Kluch")
		one8 := r.FormValue("insert_Aparatnaa")
		one9 := r.FormValue("insert_AD")
		one10 := r.FormValue("insert_Contact")
		one11 := r.FormValue("insert_WGS")
		one12 := r.FormValue("insert_Primechanie")
		fmt.Println("Xaxaxa")
		pg_db := db.Connect()
		SavePL2(pg_db, one1, one2, one3, one4, one5, one6, one7, one8, one9, one10, one11, one12)
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

func SavePL2(dbRef *pg.DB, one1 string, one2 string, one3 string, one4 string, one5 string, one6 string, one7 string, one8 string, one9 string, one10 string, one11 string, one12 string) {

	newPL1 := &db.Dostup101{
		ID:             one1,
		Diapazon:       one2,
		Adres:          one3,
		Infrastructura: one4,
		Prinadlegnost:  one5,
		Dostup:         one6,
		Kluch:          one7,
		Aparatnaa:      one8,
		AD:             one9,
		Contact:        one10,
		WGS:            one11,
		Primechanie:    one12,
		//IsActive:       true,
	}
	newPL1.Save(dbRef)
}

//type Front struct {
//	ID             string
//	Diapazon       string
//	Adres          string
//	Infrastructura string
//	Prinadlegnost  string
//	Dostup         string
//	Kluch          string
//	Aparatnaa      string
//	AD             string
//	Contact        string
//	WGS            string
//	Primechanie    string
//}
