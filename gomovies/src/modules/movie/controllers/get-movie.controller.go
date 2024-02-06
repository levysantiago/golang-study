package movies_controllers

import (
	"encoding/json"
	"fmt"
	movies_repository "gomovies/src/modules/movie/repositories"
	movies_services "gomovies/src/modules/movie/services"
	"net/http"

	"github.com/gorilla/mux"
)

type GetMovieResponseData struct{
	Status string `json:"status"`
	Data movies_repository.Movie `json:"data"`
}

type GetMovieResponseError struct{
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func GetMovie(res http.ResponseWriter, req *http.Request){
	var id = mux.Vars(req)["id"]

	var movie = movies_services.GetMovie(id)
	var isMovieInvalid bool = movie.Id == ""

	var responseData GetMovieResponseData
	if !isMovieInvalid{
		responseData.Status = "success"
		responseData.Data = movie

		jsonResponse, jsonError := json.Marshal(responseData)

		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
			var responseError GetMovieResponseError
			responseError.Status = "error"
			responseError.ErrorMessage = "Internal server error"
			
			jsonResponse, _ := json.Marshal(responseError)
			
			res.WriteHeader(http.StatusInternalServerError)
			res.Write(jsonResponse)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write(jsonResponse)
		return
	}else{
		var responseError GetMovieResponseError
		responseError.Status = "error"
		responseError.ErrorMessage = "Movie not found"
		
		jsonResponse, _ := json.Marshal(responseError)
		
		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonResponse)
		return
	}
}