package user

import (
	"context"
	"errors"
	userpb "practice/IMDB/gunk/v1/user"
	"practice/IMDB/usermgm/storage"
)

func (s *Svc) AddMovieRating(ctx context.Context, r *userpb.AddMovieRatingRequest) (*userpb.AddMovieRatingResponse, error) {
	movieRating := storage.MovieRating{
		MovieID: int(r.GetMovieID()),
		UserID:  int(r.GetUserID()),
		Rating:  int(r.GetRating()),
	}

	if err := movieRating.Validate(); err != nil {
		return nil, err
	}

	am, err := s.core.AddMovieRating(movieRating)
	if err != nil {
		return nil, err
	}
	return &userpb.AddMovieRatingResponse{
		AddMovieRating: &userpb.AddMovieRating{
			UserID:  int32(am.UserID),
			MovieID: int32(am.MovieID),
			Rating:  int32(am.Rating),
		},
	}, nil
}
func (s *Svc) EditMovieRating(ctx context.Context, r *userpb.EditMovieRatingRequest) (*userpb.EditMovieRatingResponse, error) {
	movieRating := storage.MovieRating{
		MovieID: int(r.GetMovieID()),
		UserID:  int(r.GetUserID()),
		Rating:  int(r.GetRating()),
	}

	if err := movieRating.Validate(); err != nil {
		return nil, err
	}

	am, err := s.core.EditMovieRating(movieRating)
	if err != nil {
		return nil, err
	}
	return &userpb.EditMovieRatingResponse{
		EditMovieRating: &userpb.EditMovieRating{
			UserID:  int32(am.UserID),
			MovieID: int32(am.MovieID),
			Rating:  int32(am.Rating),
		},
	}, nil
}

func (s *Svc) AddInWatchList(ctx context.Context, r *userpb.AddInWatchListRequest) (*userpb.AddInWatchListResponse, error) {
	movieWatched := storage.MovieWatched{
		MovieIDs: r.GetMovieID(),
		UserID:  int(r.GetUserID()),
	}

	if err := movieWatched.Validate(); err != nil {
		return nil, err
	}

	am, err := s.core.AddInWatchList(movieWatched)
	if err != nil {
		return nil, err
	}

	if len(am) == 0 {
		return nil, errors.New("failed to add movies")
	}

	var watchList []*userpb.AddInWatchList
	for _, k := range am {
		watchList = append(watchList, &userpb.AddInWatchList{
			UserID:  int32(k.UserID),
			MovieID: k.MovieID,
		})
	}
	
	return &userpb.AddInWatchListResponse{
		AddInWatchList: watchList,
	}, nil
}
