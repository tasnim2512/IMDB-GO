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

