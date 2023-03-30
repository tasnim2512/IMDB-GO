package postgres

import (
	"practice/IMDB/usermgm/storage"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "CREATE_USER_SUCEESS",
			in: storage.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "test@example.com",
				UserName:  "test",
				Password:  "12345678",
			},
			want: &storage.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "test@example.com",
				UserName:  "test",
				IsActive:    true,
			},
		},
		{
			name: "CREATE_USER_EMIAL_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "first2",
				LastName:  "last2",
				Email:     "test@example.com",
				UserName:  "test2",
				Password:  "12345678",
			},
			wantErr: true,
		},
		{
			name: "CREATE_USER_USRNAME_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "first3",
				LastName:  "last3",
				Email:     "test3@example.com",
				UserName:  "test",
				Password:  "12345678",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UserRegistration(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UserRegistration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "Role","IsAdmin","CreatedAt","UpdatedAt", "DeletedAt",),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UserRegistration() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}


func TestUpdateUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "tasnim",
		LastName:  "hossain",
		Email:     "tasnim@yahoo.com",
		UserName:  "prapty",
		Password:  "12345678",
	}
	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "UPDATE_USER_SUCEESS",
			in: storage.User{
				FirstName: "Tas",
				LastName:  "Nim",
			},
			want: &storage.User{
				FirstName: "Tas",
				LastName:  "Nim",
				Email:     "tasnim@yahoo.com",
				UserName:  "prapty",
				IsActive:   false,
			},
		},
		{
			name: "UPDATE_USER_Email_Error",
			in: storage.User{
				FirstName: "Tas",
				LastName:  "Nim",
				Email:     "tasnim@hhhh.com",
				UserName:  "prapty",
				IsActive:   false,
			},
			want: nil,
			wantErr: true,
		},
	}

	user, err := s.UserRegistration(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				tt.in.ID = user.ID
			}
			got, err := s.UpdateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "Role","IsAdmin","CreatedAt","UpdatedAt", "DeletedAt",),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "tasnim",
		LastName:  "hossain",
		Email:     "tasnimgmail.com",
		UserName:  "prapty",
		Password:  "12345678",
	}

	user, err := s.UserRegistration(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.UserRegistration() error = %v", err)
	}
	id := strconv.Itoa(user.ID)
	tests := []struct {
		name    string
		in      string
		wantErr bool
	}{
		{
			name: "DELETE_USER_BY_ID_SUCEESS",
			in:   id,
		},
		// {
		// 	name:    "DELETE_USER_BY_ID_FAILED",
		// 	in:      id,
		// 	wantErr: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

