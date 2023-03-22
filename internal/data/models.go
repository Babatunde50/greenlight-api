package data

import (
	"database/sql"
	"errors"
	"time"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when // looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this, // like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
		GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
	}
	Users interface {
		GetByEmail(email string) (*User, error)
		GetForToken(tokenScope string, tokenPlaintext string) (*User, error)
		Insert(user *User) error
		Update(user *User) error
	}
	Tokens interface {
		DeleteAllForUser(scope string, userID int64) error
		Insert(token *Token) error
		New(userID int64, ttl time.Duration, scope string) (*Token, error)
	}
	Permissions interface {
		AddForUser(userID int64, codes ...string) error
		GetAllForUser(userID int64) (Permissions, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Movies:      MockMovieModel{},
		Users:       MockUserModel{},
		Tokens:      MockTokenModel{},
		Permissions: MockPermissionModel{},
	}
}
