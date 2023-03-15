package admin

import (
	"context"
	adminpb "practice/IMDB/gunk/v1/admin"
	"practice/IMDB/usermgm/storage"
)

type AdminStore interface {
	Register(a storage.Admin) (*storage.Admin, error)
}

type Svc struct {
	adminpb.UnimplementedAdminServiceServer
	core AdminStore
}

func NewUserSvc(as AdminStore) *Svc {
	return &Svc{
		core: as,
	}
}

// func (s Svc) Register(a storage.Admin) (*storage.Admin, error) {

// 	admin, err := s.NewCoreUser.Register(a)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if admin == nil{
// 		log.Println("unable to register")
// 	}
// 	return admin, nil
// }

func (s *Svc) Register(ctx context.Context, r *adminpb.RegisterRequest) (*adminpb.RegisterResponse, error) {
	admin := storage.Admin{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		UserName:  r.UserName,
		Password:  r.Password,
	}
	if err := admin.Validate(); err != nil {
		return nil, err
	}
	a, err := s.core.Register(admin)
	if err != nil {
		return nil, err
	}
	return &adminpb.RegisterResponse{
		Admin: &adminpb.Admin{
			FirstName: a.FirstName,
			LastName:  a.LastName,
			UserName:  a.UserName,
			Email:     a.Email,
			IsAdmin:   a.IsAdmin,
			Role:      a.Role,
		},
	}, nil
}
