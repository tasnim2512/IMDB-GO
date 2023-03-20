package admin

import "practice/IMDB/usermgm/storage"

func (s Svc) AddGenre(g storage.Genre) (*storage.Genre, error) {
	ag, err := s.store.AddGenre(g)
	if err != nil {
		return nil, err
	}
	return ag,nil
}
