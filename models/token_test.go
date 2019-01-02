package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSessionHeader(t *testing.T) {
	t.Run("Invalid sesssion xxxxx.", func(t *testing.T) {
		vals, err := ParseToken("xxxxx.")
		assert.Empty(t, vals)
		assert.Equal(t, ErrNotValidSessionToken, err)
	})

	t.Run("Valid session xx.supertest", func(t *testing.T) {
		vals, err := ParseToken("xx.supertest")
		assert.Nil(t, err)
		assert.Equal(t, "xx", vals.Selector)
		assert.Equal(t, "supertest", vals.Validator)
	})
}
