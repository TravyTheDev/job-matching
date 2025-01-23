package users

import (
	"database/sql"

	"github.com/TravyTheDev/job-matching/service/types"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// todo return user
func (s *UserStore) CreateUser(user types.User) error {
	stmt := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)`

	_, err := s.db.Exec(stmt, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}
