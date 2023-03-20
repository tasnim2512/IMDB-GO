package admin

import "practice/IMDB/usermgm/storage"

type AdminStore interface {
	AddGenre(storage.Genre) (*storage.Genre, error)
}

type Svc struct {
	store AdminStore
}

func NewCoreAdmin(as AdminStore) *Svc {
	return &Svc{
		store: as,
	}
}

func (s Svc) AddGenre(g storage.Genre) (*storage.Genre, error) {
	ag, err := s.store.AddGenre(g)
	if err != nil {
		return nil, err
	}
	return ag,nil
}
