package data

import (
	"time"
)

var mockMovie = Movie{
	ID:        1,
	CreatedAt: time.Now(),
	Title:     "Mock Movie",
	Year:      2022,
	Runtime:   Runtime(30),
	Genres:    []string{"comedy", "action"},
	Version:   1,
}

type MockMovieModel struct{}

func (m MockMovieModel) Insert(movie *Movie) error {
	return nil
}
func (m MockMovieModel) Get(id int64) (*Movie, error) {
	return &mockMovie, nil
}
func (m MockMovieModel) Update(movie *Movie) error {
	return nil
}
func (m MockMovieModel) Delete(id int64) error {
	return nil
}
func (m MockMovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error) {
	return []*Movie{
		&mockMovie,
	}, Metadata{}, nil
}

var mockUser = User{
	ID:        1,
	CreatedAt: time.Now(),
	Name:      "Tunde Ola",
	Email:     "baba@ola.com",
	// Password:  "438953kjdsfks",
	Activated: true,
	Version:   1,
}

type MockUserModel struct{}

func (m MockUserModel) GetByEmail(email string) (*User, error) {
	return &mockUser, nil
}
func (m MockUserModel) GetForToken(tokenScope string, tokenPlaintext string) (*User, error) {
	return &mockUser, nil
}
func (m MockUserModel) Insert(user *User) error {
	return nil
}
func (m MockUserModel) Update(user *User) error {
	return nil
}

var mockToken = Token{
	Plaintext: "3945345",
	Hash:      []byte("sdlfsdjlfsdfaf"),
	UserID:    1,
	Expiry:    time.Now(),
	Scope:     "activation",
}

type MockTokenModel struct{}

func (m MockTokenModel) DeleteAllForUser(scope string, userID int64) error {
	return nil
}
func (m MockTokenModel) Insert(token *Token) error {
	return nil
}
func (m MockTokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
	return &mockToken, nil
}

var mockPermission = Permissions{"movies:read", "movies:write"}

type MockPermissionModel struct{}

func (m MockPermissionModel) AddForUser(userID int64, codes ...string) error {
	return nil
}

func (m MockPermissionModel) GetAllForUser(userID int64) (Permissions, error) {
	return mockPermission, nil
}
