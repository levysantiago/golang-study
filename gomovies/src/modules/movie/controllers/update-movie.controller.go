package movies_controllers

import (
	"encoding/json"
	movies_repository "gomovies/src/modules/movie/repositories"
	movies_services "gomovies/src/modules/movie/services"
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
	var params movies_repository.UpdateMovieDTO
	
	var id = mux.Vars(req)["id"]

	decodingError := json.NewDecoder(req.Body).Decode(&params)
	if decodingError != nil {
			http.Error(res, decodingError.Error(), http.StatusBadRequest)
			return
	}

	var movie, updateError = movies_services.UpdateMovie(id, &params)

	if(updateError != nil){
		var responseError CreateMovieResponseError
		responseError.ErrorMessage = updateError.Error()
		responseError.Status = "error"

		jsonResponse, _ := json.Marshal(responseError)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonResponse)
		return
	}

	var responseData UpdateMovieResponseData
	responseData.Status = "success"
	responseData.Data = movie

	var jsonResponse, jsonError = json.Marshal(responseData)

	if(jsonError != nil){
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