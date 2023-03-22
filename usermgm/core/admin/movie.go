package admin

import (
	"fmt"
	"practice/IMDB/usermgm/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Svc) AddMovie(m storage.Movie) (*storage.Movie, error) {

	newMovie, err := s.store.AddMovie(m)
	if err != nil {
		return nil, err
	}

	for _, val := range m.Genre {
		GenreExists, _ := s.GenreExists( int(val))
		if !GenreExists {
			return nil, status.Error(codes.AlreadyExists, "genre does not exists")
		}
		movieGenre, _ := s.store.AddMovieGenre(storage.MovieGenre{
			MovieID: newMovie.ID,
			GenreID: int(val),
		})
		fmt.Println(movieGenre)
	}
	return newMovie, nil
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
