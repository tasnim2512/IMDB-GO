package postgres

import (
	"fmt"
	"log"
	"practice/IMDB/usermgm/storage"
)

const addGenreQuery = `
INSERT INTO genres(
	name
) VALUES(
	:name
) RETURNING *;
`

func (s PostgresStorage) AddGenre(genre storage.Genre) (*storage.Genre, error) {
	stmt, err := s.DB.PrepareNamed(addGenreQuery)
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Get(&genre, genre); err != nil {
		return nil, err
	}
	if genre.ID == 0 {
		log.Println("unable to insert user")
		return nil, fmt.Errorf("unable to insert user")
	}
	return &genre, nil
}

const EditGenreQuery = `
	UPDATE genres SET
	name=:name
	WHERE id=:id AND deleted_at IS NULL
	RETURNING *;
	`

func (s PostgresStorage) EditGenre(u storage.Genre) (*storage.Genre, error) {
	stmt, err := s.DB.PrepareNamed(EditGenreQuery)
	if err != nil {
		log.Fatal(err)
	}
	if err := stmt.Get(&u, u); err != nil {
		log.Println(err)
		return nil, err
	}
	
	return &u, nil
}

const DeleteGenreQuery = `
UPDATE genres SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL 
RETURNING id;
`

func (s PostgresStorage) DeleteGenre(id string) error {
	res, err := s.DB.Exec(DeleteGenreQuery, id)
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
const GetGenreByNameQuery = `
SELECT id, name FROM genres WHERE name= $1 AND deleted_at IS NULL;
`

func (s PostgresStorage) GetGenreByName(name string) (*storage.Genre, error) {
	var genres storage.Genre
	if err := s.DB.Get(&genres, GetGenreByNameQuery, name); err != nil {
		return nil, err
	}
	if genres.ID == 0 {
		return nil, fmt.Errorf("unable to get genre")
	}
	return &genres, nil
}