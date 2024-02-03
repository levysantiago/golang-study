package movies_services

import movies_repository "gomovies/src/modules/movie/repositories"

func UpdateMovie(id string, data *movies_repository.UpdateMovieDTO) movies_repository.Movie{
	return movies_repository.Update(id, data)
}