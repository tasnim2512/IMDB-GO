package postgres

import (
	"database/sql"
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

const getMovieGenreQuery = `
SELECT genre_id FROM movie_genre WHERE movie_id = $1 AND deleted_at IS NULL;;
`

func (s PostgresStorage) GetAllMovieGenreByMovieID(id int) ([]*storage.MovieGenre, error) {
	var movie_genre []*storage.MovieGenre
	if err := s.DB.Select(&movie_genre, getMovieGenreQuery, id); err != nil {

		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}
		return nil, err
	}
	return movie_genre, nil
}

const editMovieGenreQuery = `
INSERT INTO movie_genre(
	movie_id,
	genre_id
) VALUES(
	:movie_id,
	:genre_id
) RETURNING *;
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
