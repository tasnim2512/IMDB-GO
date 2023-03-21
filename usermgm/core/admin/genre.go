package admin

import "practice/IMDB/usermgm/storage"

func (s Svc) AddGenre(g storage.Genre) (*storage.Genre, error) {
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
		return  err
	}
	return nil
}
