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

const EditMovieQuery = `
	UPDATE movies SET
	name=:name,
	storyline=:storyline
	WHERE id=:id AND deleted_at IS NULL
	RETURNING *;
	`

func (s PostgresStorage) EditMovie(m storage.Movie) (*storage.Movie, error) {
	stmt, err := s.DB.PrepareNamed(EditMovieQuery)
	if err != nil {
		log.Fatal(err)
	}
	if err := stmt.Get(&m, m); err != nil {
		log.Println(err)
		return nil, err
	}

	return &m, nil
}

const DeleteMovieQuery = `
UPDATE movies SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL 
RETURNING id;
`

func (s PostgresStorage) DeleteMovie(id string) error {
	res, err := s.DB.Exec(DeleteMovieQuery, id)
	if err != nil {
		log.Println(err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("%v", err)
	}
	if row >= 0 {
		fmt.Printf("unable to delete movie")
	}
	return nil
}

const GetMovieByIdQuery = `
SELECT id, name FROM movies WHERE id = $1 AND deleted_at IS NULL;
`

func (s PostgresStorage) GetMovieByName(name string) (*storage.Movie, error) {
	var movies storage.Movie
	if err := s.DB.Get(&movies, GetMovieByIdQuery, name); err != nil {
		return nil, err
	}
	if movies.ID == 0 {
		return nil, fmt.Errorf("unable to get movie")
	}
	return &movies, nil
}

const movieRatingQuery = `
INSERT INTO movie_rating(
	movie_id,
	user_id,
	rating
) VALUES(
	:movie_id,
	:user_id,
	:rating
) RETURNING *;
`

func (s PostgresStorage) AddMovieRating(rating storage.MovieRating) (*storage.MovieRating, error) {
	stmt, err := s.DB.PrepareNamed(movieRatingQuery)
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Get(&rating, rating); err != nil {
		return nil, err
	}
	if rating.ID == 0 {
		log.Println("unable to rate movie")
		return nil, fmt.Errorf("unable to rate movie")
	}
	return &rating, nil
}

const updateMovieRatingQuery = `
UPDATE movie_rating SET
	rating=:rating
	WHERE movie_id=:movie_id AND user_id=:user_id 
	RETURNING *;`

func (s PostgresStorage) EditMovieRating(rating storage.MovieRating) (*storage.MovieRating, error) {
	stmt, err := s.DB.PrepareNamed(updateMovieRatingQuery)
	if err != nil {
		log.Fatal(err)
	}
	if err := stmt.Get(&rating, rating); err != nil {
		log.Println(err)
		return nil, err
	}

	return &rating, nil
}


const addInWatchList = `
INSERT INTO movie_watched(
	movie_id,
	user_id
) VALUES(
	:movie_id,
	:user_id
) RETURNING *;
`

func (s PostgresStorage) AddInWatchList(watched storage.MovieWatched) (*storage.MovieWatched, error) {
	stmt, err := s.DB.PrepareNamed(addInWatchList)
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Get(&watched, watched); err != nil {
		return nil, err
	}
	if watched.ID == 0 {
		log.Println("unable to add movie in watch list")
		return nil, fmt.Errorf("unable to add movie in watch list")
	}
	return &watched, nil
}
