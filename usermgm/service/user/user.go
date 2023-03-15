package user

import (
	"context"
	userpb "practice/IMDB/gunk/v1/user"
	"practice/IMDB/usermgm/storage"
)

type UserStore interface {
	Register(storage.User) (*storage.User, error)
	Login(storage.Login) (*storage.User, error)
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
			FirstName: ls.FirstName,
			LastName:  ls.LastName,
			Email:     ls.Email,
			UserName:  ls.UserName,
			Role:      ls.Role,
			IsAdmin:   false,
			IsActive:  false,
		},
	},nil
}
