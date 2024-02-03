package movies_controllers

import (
	"encoding/json"
	movies_repository "gomovies/src/modules/movie/repositories"
	movies_services "gomovies/src/modules/movie/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateMovie(res http.ResponseWriter, req *http.Request){
	var data movies_repository.UpdateMovieDTO
	
	var id = mux.Vars(req)["id"]

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
	}

	var movie = movies_services.UpdateMovie(id, &data)

	var jsonResponse, jsonError = json.Marshal(movie)

	if jsonError != nil{
		log.Fatal("Unable to encode JSON")
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonResponse)
}