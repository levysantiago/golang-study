package routes

import (
	services "movies-crud/controllers"

	"github.com/gorilla/mux"
)

func MoviesRoutes(router *mux.Router){
	router.HandleFunc("/movies", services.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", services.GetMovie).Methods("GET")
	router.HandleFunc("/movies", services.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", services.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", services.DeleteMovie).Methods("DELETE")
}