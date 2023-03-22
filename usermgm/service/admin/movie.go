package admin

import (
	"context"
	adminpb "practice/IMDB/gunk/v1/admin"
	"practice/IMDB/usermgm/storage"
)

func (s *Svc) AddMovie(ctx context.Context, r *adminpb.AddMovieRequest) (*adminpb.AddMovieResponse, error) {
	movie := storage.Movie{
		Name:       r.GetName(),
		StoryLine:  r.GetStoryLine(),
		Genre:      r.Genres,
	}

	if err := movie.Validate(); err != nil {
		return nil, err
	}

	am, err := s.core.AddMovie(movie)
	if err != nil {
		return nil, err
	}
	return &adminpb.AddMovieResponse{
		AddMovie: &adminpb.AddMovie{
			ID:        int32(am.ID),
			Name:      am.Name,
			StoryLine: am.StoryLine,
			Genres:    am.Genre,
		},
	}, nil
}

