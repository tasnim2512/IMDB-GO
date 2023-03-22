package admin

import (
	adminpb "practice/IMDB/gunk/v1/admin"
	"practice/IMDB/usermgm/storage"
)

type CoreAdmin interface {
	AddGenre(storage.Genre) (*storage.Genre, error)
	EditGenre(storage.Genre) (*storage.Genre, error)
	DeleteGenre(string) error
	AddMovie(storage.Movie) (*storage.Movie, error)
}

type Svc struct {
	adminpb.UnimplementedAdminServiceServer
	core CoreAdmin
}

func NewAdminSvc(ua CoreAdmin) *Svc {
	return &Svc{
		core: ua,
	}
}
