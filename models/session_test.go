package models

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/haskaalo/intribox/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetSessionBySelector(t *testing.T) {
	t.Run("Should succcessfully return a session based on a selector", func(t *testing.T) {
		selector, validator, err := InitiateSession(420)
		assert.NoError(t, err)

		sess, err := GetSessionBySelector(selector)
		assert.NoError(t, err)
		assert.Equal(t, 420, sess.UserID, "UserID in redis should match what is expected")
		assert.Equal(t, selector, sess.Selector)
		assert.Equal(t, utils.SHA1([]byte(validator)), sess.Validator)
	})

	t.Run("Should return an error if the selector doesn't exist", func(t *testing.T) {
		sess, err := GetSessionBySelector("not-valid")
		assert.Nil(t, sess, "Returned Session should be nil")
		assert.EqualError(t, err, redis.Nil.Error())
	})
}

func TestGetSessionByToken(t *testing.T) {
	t.Run("Should successfully return a session based on a token", func(t *testing.T) {
		selector, validator, err := InitiateSession(420)
		assert.NoError(t, err)

		token := selector + "." + validator

		sess, err := GetSessionByToken(token)
		assert.NoError(t, err)
		assert.Equal(t, 420, sess.UserID, "UserID in redis should match what is expected")
		assert.Equal(t, selector, sess.Selector)
		assert.Equal(t, utils.SHA1([]byte(validator)), sess.Validator)
	})

	t.Run("Should return an error if token is not in a valid format", func(t *testing.T) {
		sess, err := GetSessionByToken("invalid-format-of-a-token")
		assert.Nil(t, sess, "Session returned should be nil")
		assert.EqualError(t, err, ErrNotValidSessionToken.Error())
	})

	t.Run("Should return an error if selector doesn't exist", func(t *testing.T) {
		sess, err := GetSessionByToken("abc.dfg")
		assert.EqualError(t, err, ErrNotValidSessionToken.Error())
		assert.Nil(t, sess, "Session returned should be nil")
	})

	t.Run("Should return an error if selector exist but the validator is not valid", func(t *testing.T) {
		selector, _, err := InitiateSession(420)
		assert.NoError(t, err)

		token := selector + "." + "invalid"
		sess, err := GetSessionByToken(token)

		assert.EqualError(t, err, ErrNotValidSessionToken.Error())
		assert.Nil(t, sess, "Session returned should be nil")
	})
}

func TestDeleteSessionBySelector(t *testing.T) {
	t.Run("Should a session when called", func(t *testing.T) {
		selector, _, err := InitiateSession(420)
		assert.NoError(t, err)

		err = DeleteSessionBySelector(selector)
		assert.NoError(t, err, "Should successfully terminate")

		// Check if it no longer exist in database
		_, err = GetSessionBySelector(selector)
		assert.EqualError(t, err, redis.Nil.Error())
	})
}

func TestResetTimeSession(t *testing.T) {
	t.Run("Should successfully reset the expire timer of a session", func(t *testing.T) {
		selector, _, err := InitiateSession(420)
		assert.NoError(t, err)

		sess, err := GetSessionBySelector(selector)
		assert.NoError(t, err)

		err = sess.ResetTimeSession()
		assert.NoError(t, err, "Should have no error when successfully resetting the expire timer of a session")
	})
}

func TestInitiateSession(t *testing.T) {
	t.Run("Should successfully initiate a session", func(t *testing.T) {
		selector, _, err := InitiateSession(420)
		assert.NoError(t, err, "Calling InitiateSession should have no error")

		// Check if it actually exist
		_, err = GetSessionBySelector(selector)
		assert.NoError(t, err)
	})
}

func TestDeleteOldestSession(t *testing.T) {
	t.Run("Should successfully delete the oldest session", func(t *testing.T) {
		err := r.FlushAll().Err()

		oldestSessionSelector := ""
		for i := 0; i < 5; i++ {
			selector, _, err := InitiateSession(42069)
			if i == 0 {
				oldestSessionSelector = selector
			}

			assert.NoError(t, err)
			time.Sleep(2 * time.Second) // Make sure we get the oldest session
		}

		err = DeleteOldestSession(42069)
		assert.NoError(t, err, "Calling DeleteOldestSession should have no error")

		_, err = GetSessionBySelector(oldestSessionSelector)
		assert.EqualError(t, err, redis.Nil.Error(), "The oldest session should no longer exist")
	})
}
