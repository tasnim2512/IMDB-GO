package admin

import "practice/IMDB/usermgm/storage"

type AdminStore interface {
	AddGenre(storage.Genre) (*storage.Genre, error)
	EditGenre(storage.Genre) (*storage.Genre, error)
	DeleteGenre(id string)  error
}

type Svc struct {
	store AdminStore
}

func NewCoreAdmin(as AdminStore) *Svc {
	return &Svc{
		store: as,
	}
}
