package movies_services

import (
	movies_repository "gomovies/src/modules/movie/repositories"

	"github.com/google/uuid"
)

type CreateMovieServiceDTO struct{
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director_Id string `json:"director_id"`
}

func CreateMovieService(params *CreateMovieServiceDTO) movies_repository.Movie{
	id := uuid.New().String()


	var newMovie = movies_repository.Movie{
		Id: id, 
		Isbn: params.Isbn,
		Title: params.Title,
		Director_Id: params.Director_Id,
	}
	
	movies_repository.Create(newMovie)

	return newMovie
}