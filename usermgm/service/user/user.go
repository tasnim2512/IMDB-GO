package user

import (
	"context"
	userpb "practice/IMDB/gunk/v1/user"
	"practice/IMDB/usermgm/storage"
	"strconv"
)

type UserStore interface {
	Register(storage.User) (*storage.User, error)
	UpdateUser(storage.User) (*storage.User, error)
	DeleteUser(string) error
	Login(storage.Login) (*storage.User, error)
	AddMovieRating(storage.MovieRating) (*storage.MovieRating, error)
	EditMovieRating(storage.MovieRating) (*storage.MovieRating, error)
	AddInWatchList(storage.MovieWatched) ([]*storage.MovieWatched, error)
}

type Svc struct {
	userpb.UnimplementedUserServiceServer
	core UserStore
}

func NewUserSvc(us UserStore) *Svc {
	return &Svc{
		core: us,
	}
}

func (s *Svc) Register(ctx context.Context, r *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := storage.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		UserName:  r.UserName,
		Password:  r.Password,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	a, err := s.core.Register(user)
	if err != nil {
		return nil, err
	}
	return &userpb.RegisterResponse{
		User: &userpb.User{
			ID:        int32(a.ID),
			FirstName: a.FirstName,
			LastName:  a.LastName,
			Email:     a.Email,
			UserName:  a.UserName,
			Role:      a.Role,
			IsAdmin:   false,
			IsActive:  false,
		},
	}, nil
}

func (s *Svc) Login(ctx context.Context, r *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	login := storage.Login{
		UserName: r.GetUserName(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		return nil, err
	}

	ls, err := s.core.Login(login)
	if err != nil {
		return nil, err
	}
	return &userpb.LoginResponse{
		User: &userpb.User{
			ID:        int32(ls.ID),
			FirstName: ls.FirstName,
			LastName:  ls.LastName,
			Email:     ls.Email,
			UserName:  ls.UserName,
			Role:      ls.Role,
			IsAdmin:   false,
			IsActive:  false,
		},
	}, nil
}

func (s *Svc) UpdateUser(ctx context.Context, r *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user := storage.User{
		ID:        int(r.GetID()),
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Role:      r.GetRole(),
		IsActive:  r.GetIsActive(),
	}
	a, err := s.core.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{
		UpdateUser: &userpb.UpdateUser{
			FirstName: a.FirstName,
			LastName:  a.LastName,
			Role:      a.Role,
			IsActive:  false,
		},
	}, nil

}

func (s *Svc) DeleteUser(ctx context.Context, r *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	genre := storage.Genre{
		ID: int(r.GetID()),
	}

	gID := strconv.Itoa(genre.ID)
	_ = s.core.DeleteUser(gID)

	return &userpb.DeleteUserResponse{
		Error: "Deleted",
	}, nil
}
