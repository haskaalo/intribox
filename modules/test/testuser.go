package test

import (
	"github.com/haskaalo/intribox/models"
)

// TestUserPassword Password used for CreateTestUser
var TestUserPassword = "137945"

// CreateTestUser Create a test user
func CreateTestUser() (*models.User, error) {
	user := &models.User{
		Email:    "test@example.com",
		Password: []byte("$2y$12$LXAwwYDwaHY7dR/LM8QzIOWE.nqbJ7wor/u7KZBrh3e6wnlqZsn66"),
	}

	err := user.InsertNewUser()
	if err != nil {
		return nil, err
	}

	usr, err := models.GetUserByEmail("test@example.com")
	return usr, err
}
