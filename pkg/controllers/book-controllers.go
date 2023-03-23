package main

import (
	"log"
	"net/http"

	_ "gihtub.com/jinzhu/gorm/dialects/mysql"
	"github.com/anassidr/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
