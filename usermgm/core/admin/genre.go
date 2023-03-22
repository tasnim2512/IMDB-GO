package admin

import (
	"practice/IMDB/usermgm/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Svc) AddGenre(g storage.Genre) (*storage.Genre, error) {
	alreadyExists, _ := s.GenreAlreadyExists(g.Name)

	if alreadyExists {
		return nil, status.Error(codes.AlreadyExists, "name already exists")
	}
	ag, err := s.store.AddGenre(g)
	if err != nil {
		return nil, err
	}
	return ag, nil
}

func (s Svc) EditGenre(g storage.Genre) (*storage.Genre, error) {
	ag, err := s.store.EditGenre(g)
	if err != nil {
		return nil, err
	}
	return ag, nil
}

func (s Svc) DeleteGenre(g string) error {
	err := s.store.DeleteGenre(g)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) GenreAlreadyExists(value string) (bool, error) {
	newGenre, err := s.store.GetGenreByName(value)
	if err != nil {
		return false, err
	}
	if newGenre != nil {
		return true, nil
	}
	return false, nil
}
