package admin

import "practice/IMDB/usermgm/storage"

func (s Svc) EditGenre(g storage.Genre) (*storage.Genre, error) {
	ag, err := s.store.EditGenre(g)
	if err != nil {
		return nil, err
	}
	return ag, nil
}
