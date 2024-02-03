package routes

import (
	movies_controllers "movies-crud/src/modules/movie/controllers"

	"github.com/gorilla/mux"
)

func MoviesRoutes(router *mux.Router){
	router.HandleFunc("/movies", movies_controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", movies_controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies", movies_controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", movies_controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", movies_controllers.DeleteMovie).Methods("DELETE")
}