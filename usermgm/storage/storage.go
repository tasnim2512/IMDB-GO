package storage

import (
	"database/sql"
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	NotFound = errors.New("not found")
)

type StudentFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}
type User struct {
	ID        int          `form:"-" db:"id"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	Email     string       `db:"email"`
	UserName  string       `db:"username"`
	Password  string       `db:"password"`
	Role      string       `db:"role"`
	IsAdmin   bool         `db:"is_admin"`
	IsActive  bool         `db:"is_active"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Login struct {
	UserName string `db:"username"`
	Password string `db:"password"`
}

type Genre struct {
	ID        int          `form:"-" db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Movie struct {
	ID         int          `form:"-" db:"id"`
	Name       string       `db:"name"`
	StoryLine  string       `db:"storyline"`
	Genre      []int32      `db:"genre"`
	ReleasedAt time.Time    `db:"released_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
}

type MovieGenre struct {
	ID        int          `form:"-" db:"id"`
	MovieID   int          `db:"movie_id"`
	GenreID   int          `db:"genre_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type MovieRating struct {
	ID        int          `form:"-" db:"id"`
	MovieID   int          `db:"movie_id"`
	UserID    int          `db:"user_id"`
	Rating    int          `db:"rating"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (s User) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.FirstName,
			validation.Required.Error("the First name field is required"),
		),
		validation.Field(&s.LastName,
			validation.Required.Error("the Last name field is required"),
		),
		validation.Field(&s.UserName,
			validation.Required.Error("the User name field is required"),
		),
	)
}
func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.UserName,
			validation.Required.Error("the username field is required"),
		),
		validation.Field(&l.Password,
			validation.Required.Error("the password field is required"),
		),
	)
}

func (g Genre) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Name,
			validation.Required.Error("the name field is required"),
		),
	)
}
func (m Movie) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name,
			validation.Required.Error("the name field is required"),
		),
		validation.Field(&m.StoryLine,
			validation.Required.Error("the storyline field is required"),
		),
	)
}

func (m MovieRating) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Rating,
			validation.Required.Error("the rating field is required"),
			
		),
	)
}
