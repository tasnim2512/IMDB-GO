package admin

import (
	"context"
	adminpb "practice/IMDB/gunk/v1/admin"
	"practice/IMDB/usermgm/storage"
	"strconv"
)

func (s *Svc) AddMovie(ctx context.Context, r *adminpb.AddMovieRequest) (*adminpb.AddMovieResponse, error) {
	movie := storage.Movie{
		Name:      r.GetName(),
		StoryLine: r.GetStoryLine(),
		Genre:     r.Genres,
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

func (s *Svc) EditMovie(ctx context.Context, r *adminpb.EditMovieRequest) (*adminpb.EditMovieResponse, error) {
	movie := storage.Movie{
		ID:        int(r.GetID()),
		Name:      r.GetName(),
		StoryLine: r.GetStoryLine(),
		Genre:     r.Genres,
	}
	if err := movie.Validate(); err != nil {
		return nil, err
	}

	am, err := s.core.EditMovie(movie)
	if err != nil {
		return nil, err
	}
	return &adminpb.EditMovieResponse{
		EditMovie: &adminpb.EditMovie{
			ID:        int32(am.ID),
			Name:      am.Name,
			StoryLine: am.StoryLine,
			Genres: am.Genre,
		},
	}, nil

}

func (s *Svc) DeleteMovie(ctx context.Context, r *adminpb.DeleteMovieRequest) (*adminpb.DeleteMovieResponse, error) {
	movie := storage.Movie{
		ID: int(r.GetID()),
	}

	mID := strconv.Itoa(movie.ID)
	_ = s.core.DeleteMovie(mID)

	return &adminpb.DeleteMovieResponse{
		Error: "Deleted",
	}, nil
}


