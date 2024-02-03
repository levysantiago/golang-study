package movies_controllers

import (
	"encoding/json"
	"fmt"
	movies_services "gomovies/src/modules/movie/services"
	"net/http"
)



func CreateMovie(res http.ResponseWriter, req *http.Request){
	var data movies_services.ICreateMovieServiceDTO

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
	}
	
	var movie = movies_services.CreateMovieService(&data)

	var jsonResponse, jsonError = json.Marshal(movie)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	res.WriteHeader(http.StatusCreated)
	res.Write(jsonResponse)
}