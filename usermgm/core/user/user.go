package user

import (
	"log"
	"practice/IMDB/usermgm/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	UserRegistration(storage.User) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
	UpdateUser(storage.User) (*storage.User, error)
	DeleteUser(string) error
	AddMovieRating(storage.MovieRating) (*storage.MovieRating, error)
	EditMovieRating(storage.MovieRating) (*storage.MovieRating, error)
	AddInWatchList(storage.MovieWatched) (*storage.MovieWatched, error)
	GetUserList(storage.UserFilter) ([]storage.User, error)
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

func (s Svc) Login(l storage.Login) (*storage.User, error) {
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

func (s Svc) UpdateUser(a storage.User) (*storage.User, error) {
	user, err := s.store.UpdateUser(a)
	if err != nil {
		return nil, err
	}

	if user == nil {
		log.Println("unable to register")
	}
	return user, nil
}

func (s Svc) DeleteUser(g string) error {
	err := s.store.DeleteUser(g)
	if err != nil {
		return err
	}
	return nil
}

func (s Svc) GetUserList(u storage.UserFilter) ([]storage.User, error){
	user, err := s.store.GetUserList(u)
	if err != nil {
		return nil, err
	}

	if user == nil {
		log.Println("unable to get list")
	}
	return user, nil
}
