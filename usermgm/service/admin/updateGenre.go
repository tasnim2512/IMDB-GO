package admin

import (
	"context"
	"fmt"
	adminpb "practice/IMDB/gunk/v1/admin"
	"practice/IMDB/usermgm/storage"
)

func (s *Svc) EditGenre(ctx context.Context, r *adminpb.EditGenreRequest) (*adminpb.EditGenreResponse, error) {
	genre := storage.Genre{
		ID: int(r.GetID()),
		Name: r.GetName(),
	}
	if err := genre.Validate(); err != nil {
		return nil, err
	}

	res, err := s.core.EditGenre(genre)
	
	if err != nil {
		return nil, err
	}

	return &adminpb.EditGenreResponse{
		EditGenre: &adminpb.EditGenre{
			ID:   int32(res.ID),
			Name: res.Name,
		},
	}, nil
}
