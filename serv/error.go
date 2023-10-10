package serv

import (
	"fmt"
	"html/template"
	"net/http"

	"main.go/structs"
)

func errors(w http.ResponseWriter, r *http.Request, p structs.Error) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Println("File: error.html is not found")
		return
	}
	if p.Body == "404" {
		w.WriteHeader(http.StatusNotFound)
	} else if p.Body == "500" {
		w.WriteHeader(http.StatusInternalServerError)
	} else if p.Body == "400" {
		w.WriteHeader(http.StatusBadRequest)
	} else if p.Body == "405" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	Errors = structs.Error{}
	t.Execute(w, p)
}
