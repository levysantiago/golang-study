package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lvy/gobooks/pkg/routes"
)

func main(){
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe("localhost:3333", router))
}