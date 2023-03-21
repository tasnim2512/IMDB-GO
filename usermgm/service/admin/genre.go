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
			ID:   int32(ag.ID),
			Name: ag.Name,
		},
	}, nil
}

func (s *Svc) EditGenre(ctx context.Context, r *adminpb.EditGenreRequest) (*adminpb.EditGenreResponse, error) {
	genre := storage.Genre{
		ID:   int(r.GetID()),
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

func (s *Svc) DeleteGenre(ctx context.Context, r *adminpb.DeleteGenreRequest) (*adminpb.DeleteGenreResponse, error) {
	genre := storage.Genre{
		ID:   int(r.GetID()),
	}

	res, err := s.core.EditGenre(genre)

	if err != nil {
		return nil, err
	}

	return &adminpb.DeleteGenreResponse{
		ID: int32(res.ID),
	}, nil
}
