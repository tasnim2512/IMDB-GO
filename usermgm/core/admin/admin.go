package admin

import (
	"log"
	"practice/IMDB/usermgm/storage"

	"golang.org/x/crypto/bcrypt"
)

type AdminStore interface {
	AdminRegistration(a storage.Admin) (*storage.Admin, error)
}

type Svc struct {
	store AdminStore
}

func NewCoreUser(as AdminStore) *Svc {
	return &Svc{
		store: as,
	}
}

func (s Svc) Register(a storage.Admin) (*storage.Admin, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	a.Password = string(hashPass)
	admin, err := s.store.AdminRegistration(a)
	if err != nil {
		return nil, err
	}

	if admin == nil {
		log.Println("unable to register")
	}
	return admin, nil
}
