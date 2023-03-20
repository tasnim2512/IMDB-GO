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
