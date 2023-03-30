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

const updateUserQuery = `
	UPDATE users SET
		first_name = COALESCE(NULLIF(:first_name, ''), first_name),
		last_name = COALESCE(NULLIF(:last_name, ''), last_name),
		is_active = :is_active,
		role = COALESCE(NULLIF(:role, 'user'), role)
	WHERE id = :id AND deleted_at IS NULL RETURNING *;
	`

func (s PostgresStorage) UpdateUser(u storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(updateUserQuery)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := stmt.Get(&u, u); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const deleteUserQuery = `
UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL 
RETURNING id;
	`

func (s PostgresStorage) DeleteUser(id string) error {
	res, err := s.DB.Exec(deleteUserQuery, id)
	if err != nil {
		log.Println(err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("%v", err)
	}
	if row >= 0 {
		fmt.Printf("unable to delete genre")
	}
	return nil
}

const getListQuery = `
	WITH tot AS (select count(*) as total FROM users
	WHERE
		deleted_at IS NULL
		AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%'))
	SELECT *, tot.total as total FROM users
	LEFT JOIN tot ON TRUE
	WHERE
		deleted_at IS NULL
		AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
		ORDER BY id DESC
		OFFSET $2
		LIMIT $3`

func (s PostgresStorage) GetUserList(uf storage.UserFilter) ([]storage.User, error) {
	var listUser []storage.User
	if uf.Limit == 0 {
		uf.Limit = 5
	}
	if err := s.DB.Select(&listUser, getListQuery, uf.SearchTerm, uf.Offset, uf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}
	return listUser, nil
}
