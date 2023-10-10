package main

import (
	"database/sql"

	"main.go/bd"
	"main.go/serv"
	"main.go/structs"
)

var BD *sql.DB

func main() {
	BD = bd.CreateBd()
	var errors = structs.Error{}
	if BD == nil {
		errors.Title = "500"
		errors.Body = "500"
	}
	serv.Serve(BD, errors)

}
