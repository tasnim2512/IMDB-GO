package admin

import (
	"practice/IMDB/usermgm/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Svc) AddMovie(m storage.Movie) (*storage.Movie, error) {
	newMovie, err := s.store.AddMovie(m)
	if err != nil {
		return nil, err
	}
	movieAlreadyExists, _ := s.MovieAlreadyExists(m.Name)
	if movieAlreadyExists {
		return nil, status.Error(codes.AlreadyExists, "movie already exists")
	}
	for _, val := range m.Genre {
		GenreExists, _ := s.GenreExists(int(val))
		if !GenreExists {
			return nil, status.Error(codes.AlreadyExists, "genre does not exists")
		}
		_, err := s.store.AddMovieGenre(storage.MovieGenre{
			MovieID: newMovie.ID,
			GenreID: int(val),
		})
		if err != nil {
			return nil, err
		}
	}
	return newMovie, nil
}

func (s Svc) EditMovie(m storage.Movie) (*storage.Movie, error) {
	newMovie, err := s.store.EditMovie(m)
	if err != nil {
		return nil, err
	}
	movieGenre, err := s.store.GetAllMovieGenreByMovieID(newMovie.ID)
	if err != nil {
		return nil, err
	}
	var uniqueGenreIDs []int
	for _, v1 := range newMovie.Genre {
		isUnique := true
		for _, v2 := range movieGenre {
			if int(v1) == v2.GenreID {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniqueGenreIDs = append(uniqueGenreIDs, int(v1))
		}
	}

	for _, genreID := range uniqueGenreIDs {
		GenreExists, _ := s.GenreExists(genreID)
		if !GenreExists {
			return nil, status.Error(codes.AlreadyExists, "Please enter an existing genre")
		}
		_, err := s.store.EditMovieGenre(storage.MovieGenre{
			MovieID: newMovie.ID,
			GenreID: genreID,
		})
		if err != nil {
			return nil, err
		}
	}

	return newMovie, nil
}
func (s Svc) DeleteMovie(m string) error {
	err := s.store.DeleteMovie(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) GenreExists(value int) (bool, error) {
	newGenre, err := s.store.GetGenreByID(value)
	if err != nil {
		return false, err
	}
	if newGenre != nil {
		return true, nil
	}
	return false, nil
}

func (s *Svc) MovieAlreadyExists(value string) (bool, error) {
	newGenre, err := s.store.GetMovieByName(value)
	if err != nil {
		return false, err
	}
	if newGenre != nil {
		return true, nil
	}
	return false, nil
}

