package models

import (
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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

func getUserByEmail(q sqlx.Ext, email string) (*User, error) {
	user := &User{}
	err := sqlx.Get(q, user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, knownDatabaseError(err)
	}

	return user, nil
}

// InsertNewUser Insert a new user
func (u *User) InsertNewUser() error {
	return u.insertNewUser(db)
}

func (u *User) insertNewUser(q sqlx.Ext) error {
	_, err := sqlx.NamedExec(q, "INSERT INTO users (email, password) VALUES (:email, :password)",
		map[string]interface{}{
			"email":    u.Email,
			"password": u.Password,
		})

	return knownDatabaseError(err)
}

// LogInUser Return user data if password and email match
func LogInUser(email string, password string) (*User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, knownDatabaseError(err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteAllUsers Delete all users in table user
func DeleteAllUsers() error {
	_, err := db.NamedExec("DELETE FROM users", map[string]interface{}{})
	return err
}
