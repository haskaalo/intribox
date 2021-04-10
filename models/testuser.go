package models

import "github.com/haskaalo/intribox/utils"

// TestUserPassword Password used for CreateTestUser
var TestUserPassword = "137945"

// TestingUserSession Test user session information
type TestingUserSession struct {
	Selector         string
	Validator        string
	FullSessionToken string
}

// CreateTestUser Create a test user (delete )
func CreateTestUser() (*User, error) {
	user := &User{
		Email:    utils.RandString(5) + "@example.com",
		Password: []byte("$2y$12$LXAwwYDwaHY7dR/LM8QzIOWE.nqbJ7wor/u7KZBrh3e6wnlqZsn66"), // The password is: 137945 hehe,
	}

	err := user.InsertNewUser()
	if err != nil {
		return nil, err
	}

	usr, err := GetUserByEmail(user.Email)
	return usr, err
}

// CreateTestUserSession Initiate a new session for a test user
func CreateTestUserSession(userID int) (*TestingUserSession, error) {
	selector, validator, err := InitiateSession(userID)

	s := new(TestingUserSession)
	s.Selector = selector
	s.Validator = validator
	s.FullSessionToken = selector + "." + validator

	return s, err
}
