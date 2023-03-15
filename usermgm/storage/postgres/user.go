package postgres

import (
	"fmt"
	"log"
	"practice/IMDB/usermgm/storage"
)

const userInsertQuery = `
INSERT INTO users(
	first_name,
	last_name,
	email,
	username,
	password
) VALUES(
	:first_name,
	:last_name,
	:email,
	:username,
	:password
) RETURNING *;
`

func (s PostgresStorage) UserRegistration(a storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(userInsertQuery)
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Get(&a, a); err != nil {
		return nil, err
	}
	if a.ID == 0 {
		log.Println("unable to insert user")
		return nil, fmt.Errorf("unable to insert user")
	}
	return &a, nil
}

const GetUserByUsernameQuery = `
 RETURNING *;
`

func (s PostgresStorage) GetUserByUsername(a storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(GetUserByUsernameQuery)
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Get(&a, a); err != nil {
		return nil, err
	}
	if a.ID == 0 {
		log.Println("unable to insert user")
		return nil, fmt.Errorf("unable to insert user")
	}
	return &a, nil
}
