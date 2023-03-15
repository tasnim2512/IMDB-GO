package postgres

import (
	"fmt"
	"log"
	"practice/IMDB/usermgm/storage"
)

const adminInsertQuery = `
INSERT INTO admin(
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

func (s PostgresStorage) AdminRegistration(a storage.Admin) (*storage.Admin, error) {
	stmt, err := s.DB.PrepareNamed(adminInsertQuery)
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
