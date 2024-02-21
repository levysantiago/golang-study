package movies_services

import movies_repository "gomovies/src/modules/movie/repositories"

type DeleteMovieServiceDTO struct{
	MovieId string `json:"movie_id"`
}

func DeleteMovie(params *DeleteMovieServiceDTO) (movies_repository.Movie, error){
	return movies_repository.Delete(params.MovieId)
}