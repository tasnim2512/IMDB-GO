package admin

import (
	"context"
	adminpb "practice/IMDB/gunk/v1/admin"
	"practice/IMDB/usermgm/storage"
)

func (s *Svc) AddGenre(ctx context.Context, r *adminpb.AddGenreRequest) (*adminpb.AddGenreResponse, error) {
	genre := storage.Genre{
		Name: r.GetName(),
	}
	if err := genre.Validate(); err != nil {
		return nil, err
	}

	ag, err := s.core.AddGenre(genre)
	if err != nil {
		return nil, err
	}
	return &adminpb.AddGenreResponse{
		AddGenre: &adminpb.AddGenre{
			ID:   0,
			Name: ag.Name,
		},
	}, nil
}
