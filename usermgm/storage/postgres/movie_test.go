package postgres

import (
	"practice/IMDB/usermgm/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAddMovie(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.Movie
		want    *storage.Movie
		wantErr bool
	}{
		{
			name: "CREATE_Movie_SUCEESS",
			in: storage.Movie{
				Name:      "Movie1",
				StoryLine: "MovieStory1",
			},
			want: &storage.Movie{
				Name:      "Movie1",
				StoryLine: "MovieStory1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AddMovie(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.AddMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Movie{}, "ID", "Genre", "ReleasedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.AddMovie() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateMovie(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	NewMovie := storage.Movie{
		Name:      "Movie1",
		StoryLine: "MovieStory1",
	}
	tests := []struct {
		name    string
		in      storage.Movie
		want    *storage.Movie
		wantErr bool
	}{
		{
			name: "UPDATE_Movie_SUCEESS",
			in: storage.Movie{
				Name:      "Avengers",
				StoryLine: "MovieStory1",
			},
			want: &storage.Movie{
				Name:      "Avengers",
				StoryLine: "MovieStory1",
			},
		},
	}

	movie, err := s.AddMovie(NewMovie)
	if err != nil {
		t.Fatalf("PostgresStorage.EditMovie() error = %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
				tt.in.ID = movie.ID
			got, err := s.EditMovie(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.EditMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Movie{}, "ID", "Genre", "ReleasedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.EditMovie() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
