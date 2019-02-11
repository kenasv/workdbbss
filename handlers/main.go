package handlers

import (
	"fmt"
	"html/template"
	"net/http"
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
	}

}
