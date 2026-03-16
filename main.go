package main

import (
	"net/http"

	"github.com/Wilgrey2004/go-res-api/db"
	"github.com/Wilgrey2004/go-res-api/routes"
	"github.com/gorilla/mux"
)


func main() {

	db.DBConnection()
	
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandle)

	http.ListenAndServe(":3000",r)
}