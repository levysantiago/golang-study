package movies_controllers

import (
	"encoding/json"
	movies_repository "gomovies/src/modules/movie/repositories"
	movies_services "gomovies/src/modules/movie/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UpdateMovieResponseData struct{
	Status string `json:"status"`
	Data movies_repository.Movie `json:"data"`
}

type UpdateMovieResponseError struct{
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func UpdateMovie(res http.ResponseWriter, req *http.Request){
	var data movies_repository.UpdateMovieDTO
	
	var id = mux.Vars(req)["id"]

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
	}

	var movie = movies_services.UpdateMovie(id, &data)

	var responseData UpdateMovieResponseData
	responseData.Status = "success"
	responseData.Data = movie

	var jsonResponse, jsonError = json.Marshal(responseData)

	if jsonError != nil{
		log.Fatal("Unable to encode JSON")
		
		var responseError CreateMovieResponseError
		responseError.Status = "error"
		responseError.ErrorMessage = "Internal server error"

		jsonResponse, _ = json.Marshal(responseError)

		res.Write(jsonResponse)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonResponse)
}