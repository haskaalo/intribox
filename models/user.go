package models

import (
	"time"
)

// User SQL table users
type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  []byte    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdat" db:"createdat"`
	UpdatedAt time.Time `json:"updatedat" db:"updatedat"`
	LastLogin time.Time `json:"lastlogin" db:"lastlogin"`
}

// GetUserByEmail Select user with a email
func GetUserByEmail(email string) (*User, error) {
	return getUserByEmail(db, email)
}

func getUserByEmail(q Querier, email string) (*User, error) {
	user := &User{}
	err := q.Get(user, "SELECT * FROM users WHERE email = $1", email)

	return user, err
}

// InsertNewUser Insert a new user
func (u *User) InsertNewUser() error {
	return u.insertNewUser(db)
}

func (u *User) insertNewUser(q Querier) error {
	_, err := q.NamedExec("INSERT INTO users (email, password) VALUES (:email, :password)",
		map[string]interface{}{
			"email":    u.Email,
			"password": u.Password,
		})
	return err
}

// DeleteAllUsers Delete all users in table user
func DeleteAllUsers() error {
	_, err := db.NamedExec("DELETE FROM users", map[string]interface{}{})
	return err
}
