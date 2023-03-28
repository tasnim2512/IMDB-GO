package user

import (
	"errors"
	"practice/IMDB/usermgm/storage"
)

func (s Svc) AddMovieRating(rating storage.MovieRating) (*storage.MovieRating, error) {
	newMovieRating, err := s.store.AddMovieRating(rating)
	if err != nil {
		return nil, err
	}

	return newMovieRating, nil
}

func (s Svc) EditMovieRating(rating storage.MovieRating) (*storage.MovieRating, error) {
	newMovieRating, err := s.store.EditMovieRating(rating)
	if err != nil {
		return nil, err
	}

	return newMovieRating, nil
}

func (s Svc) AddInWatchList(watched storage.MovieWatched) ([]*storage.MovieWatched, error) {
	if len(watched.MovieIDs) == 0{
		return nil, errors.New("no movie to add")
	}
	var newMovieWatched []*storage.MovieWatched
	for _ , v := range watched.MovieIDs{
		res, err := s.store.AddInWatchList(storage.MovieWatched{
			MovieID:   v,
			UserID:    watched.UserID,
		})
		if err != nil {
			return nil, err
		}
		newMovieWatched = append(newMovieWatched, res)
	}

	return newMovieWatched, nil
}
