package user

import (
	"log"
	"practice/IMDB/usermgm/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	UserRegistration(storage.User) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
}

type Svc struct {
	store UserStore
}

func NewCoreUser(as UserStore) *Svc {
	return &Svc{
		store: as,
	}
}

func (s Svc) Register(a storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	a.Password = string(hashPass)

	user, err := s.store.UserRegistration(a)
	if err != nil {
		return nil, err
	}

	if user == nil {
		log.Println("unable to register")
	}
	return user, nil
}

func (s Svc) Login(l storage.Login) (*storage.User, error){
	u, err := s.store.GetUserByUsername(l.UserName)
	if err != nil {
		return nil, err
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(l.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	l.Password = string(hashPass)

	return u, nil
}