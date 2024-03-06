package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lvy/gobooks/pkg/routes"
)

func main(){
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)

	fmt.Println("Starting server at port 3333")
	log.Fatal(http.ListenAndServe("localhost:3333", router))
}