package movies_repository

import (
	"errors"
	"log"
)

type Movie struct{
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director_Id string `json:"director_id"`
}

var movies []Movie

type UpdateMovieDTO struct{
	Isbn string `json:"isbn"`
	Title string `json:"title"`
}

func Create(newMovie Movie) {
	movies = append(movies, newMovie)
}

func Delete(id string) (Movie, error) {
	var moviesLen = len(movies)
	if(moviesLen < 1){
		err := errors.New("Movie not found")
		var emptyMovie Movie
		return emptyMovie, err
	}
	
	var index = -1
	for i:=0;i < moviesLen;i++ {
		if(movies[i].Id == id){
			index = i
		}
	}

	if(index < 0){
		err := errors.New("Movie not found")
		var emptyMovie Movie
		return emptyMovie, err
	}

	movie := movies[index]
	movies = append(movies[:index], movies[index+1:]...)

	return movie, nil
}

func FindById(id string) Movie{
	var foundMovie Movie
	
	for _, m := range movies {
		if(m.Id == id){
			foundMovie = m
		}
	}

	return foundMovie
}

func Update(id string, data *UpdateMovieDTO) Movie{
	var foundAtIndex = -1
	for i := 0; i < len(movies);i++ {
		if(movies[i].Id == id){
			foundAtIndex = i
		}
	}

	if(foundAtIndex < 0){
		log.Fatal("Movie not found")
	}

	movies[foundAtIndex].Isbn = data.Isbn
	movies[foundAtIndex].Title = data.Title

	return movies[foundAtIndex]
}

func FindMany() []Movie{
	return movies;
}