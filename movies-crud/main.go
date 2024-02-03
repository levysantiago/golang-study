package main

import (
	"fmt"
	"log"
	"movies-crud/src/modules/movie/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	routes.MoviesRoutes(router)

	fmt.Println("Starting server at port 3333")
	log.Fatal(http.ListenAndServe("localhost:3333", router))
}