package postgres

import (
	"fmt"
	"log"
	"practice/IMDB/usermgm/storage"
)

const addMovieGenreQuery = `
INSERT INTO movie_genre(
	movie_id,
	genre_id
) VALUES(
	:movie_id,
	:genre_id
) RETURNING *;
`

func (s PostgresStorage) AddMovieGenre(movieGenre storage.MovieGenre) (*storage.MovieGenre, error) {
	stmt, err := s.DB.PrepareNamed(addMovieGenreQuery)
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Get(&movieGenre, movieGenre); err != nil {
		return nil, err
	}
	if movieGenre.ID == 0 {
		return nil, fmt.Errorf("unable to insert movieGenre")
	}
	return &movieGenre, nil
}

const editMovieGenreQuery = `
UPDATE movie_genre SET
movie_id = :movie_id,
genre_id = :genre_id
WHERE id=:id AND deleted_at IS NULL
RETURNING *;
`

func (s PostgresStorage) EditMovieGenre(movieGenre storage.MovieGenre) (*storage.MovieGenre, error) {
	stmt, err := s.DB.PrepareNamed(editMovieGenreQuery)
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Get(&movieGenre, movieGenre); err != nil {
		return nil, err
	}
	if movieGenre.ID == 0 {
		return nil, fmt.Errorf("unable to edit movieGenre")
	}
	return &movieGenre, nil
}
