package postgres

import (
	"fmt"
	"log"
	"practice/IMDB/usermgm/storage"
)

const addMovieQuery = `
INSERT INTO movies(
	name,
	storyline
) VALUES(
	:name,
	:storyline
) RETURNING *;
`

func (s PostgresStorage) AddMovie(movie storage.Movie) (*storage.Movie, error) {
	stmt, err := s.DB.PrepareNamed(addMovieQuery)
	if err != nil {
		log.Fatal(err)
	}
	var ddd storage.Movie
	if err := stmt.Get(&ddd, movie); err != nil {
		return nil, err
	}
	if ddd.ID == 0 {
		return nil, fmt.Errorf("unable to insert movie")
	}
	return &ddd, nil
}
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