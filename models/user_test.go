package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByEmail(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)

	user, err := CreateTestUser()
	assert.NoError(t, err)

	t.Run("Should successfully return an valid user", func(t *testing.T) {

		userInDatabase, err := GetUserByEmail(user.Email)
		assert.NoError(t, err, "Calling GetUserByEmail should have no error")
		assert.Equal(t, user, userInDatabase, "Returned user should match what was inserted")
	})

	t.Run("Should return ErrRecordNotFound when the user doesn't exist in database", func(t *testing.T) {
		err := DeleteAllUsers()
		assert.NoError(t, err)

		userInDatabase, err := GetUserByEmail(user.Email)
		assert.EqualError(t, err, ErrRecordNotFound.Error(), "Returned error should be \"Record not found\"")
		assert.Nil(t, userInDatabase, "GetUserByEmail should return an nil user")
	})
}

func TestInsertNewUser(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)

	testUser := &User{
		Email:    "test@example.com",
		Password: []byte("some password hash"),
	}

	t.Run("Should successfully insert a valid user", func(t *testing.T) {
		err := testUser.InsertNewUser()
		assert.NoError(t, err, "Inserting user function should have no error")

		// Check if it actually exist
		userInDatabase, err := GetUserByEmail(testUser.Email)
		assert.NoError(t, err)
		assert.Equal(t, testUser.Email, userInDatabase.Email, "Email should exist in database")
		assert.Equal(t, testUser.Password, userInDatabase.Password, "Password hash should be in database")
	})
}

func TestLogInUser(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)

	t.Run("Should successfully login an user", func(t *testing.T) {
		defer func() {
			_ = DeleteAllUsers()
		}()
		testUser, err := CreateTestUser()
		assert.NoError(t, err)

		user, err := LogInUser(testUser.Email, TestUserPassword)
		assert.NoError(t, err, "Calling LogInUser should have no error")
		assert.Equal(t, user, testUser, "Returned user should match what was expected")
	})

	t.Run("Should return ErrRecordNotFound if password doesn't match but email does", func(t *testing.T) {
		defer func() {
			_ = DeleteAllUsers()
		}()
		testUser, err := CreateTestUser()
		assert.NoError(t, err)

		user, err := LogInUser(testUser.Email, "Obviously-not-the-right-password")
		assert.Nil(t, user, "Expect returned user to be nil")
		assert.EqualError(t, err, ErrRecordNotFound.Error(), "Returned error should be \"Record not found\"")
	})

	t.Run("Should return ErrRecordNotFound if email doesn't exist", func(t *testing.T) {
		defer func() {
			_ = DeleteAllUsers()
		}()
		_, err := CreateTestUser()
		assert.NoError(t, err)

		user, err := LogInUser("noexistinguser@example.com", TestUserPassword)
		assert.Nil(t, user, "Expected returned user to be nil")
		assert.EqualError(t, err, ErrRecordNotFound.Error(), "Returned error should be \"Record not found\"")
	})
}
