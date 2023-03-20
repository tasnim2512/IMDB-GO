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
 SELECT * from users WHERE username=$1 AND Deleted_at IS NULL;
`

func (s PostgresStorage) GetUserByUsername(username string) (*storage.User, error) {
	var user storage.User
	if err := s.DB.Get(&user, GetUserByUsernameQuery, username); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to get user")
	}
	return &user, nil
}
