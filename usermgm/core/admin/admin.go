package admin

import "practice/IMDB/usermgm/storage"

type AdminStore interface {
	AddGenre(storage.Genre) (*storage.Genre, error)
	EditGenre(storage.Genre) (*storage.Genre, error)
	DeleteGenre(id string) error
	GetGenreByName(string) (*storage.Genre, error)
	GetGenreByID(id int) (*storage.Genre, error)
	AddMovie(storage.Movie) (*storage.Movie, error)
	AddMovieGenre(storage.MovieGenre) (*storage.MovieGenre, error)
}

type Svc struct {
	store AdminStore
}

func NewCoreAdmin(as AdminStore) *Svc {
	return &Svc{
		store: as,
	}
}
