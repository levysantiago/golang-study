package main

import (
	"fmt"
	"log"
	routes "movies-crud/routes"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct{
	id string `json:"id"`
	isbn string `json"isbn"`
	title string `json"title"`
	director *Director `json"director"`
}

type Director struct{
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
}

var movies []Movie

func main(){
	router := mux.NewRouter()

	routes.MoviesRoutes(router)

	fmt.Println("Starting server at port 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}