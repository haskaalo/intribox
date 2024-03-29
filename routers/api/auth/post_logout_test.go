package auth

import (
	"net/http"
	"testing"

	"github.com/go-redis/redis"

	"github.com/haskaalo/intribox/middlewares"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestPostLogout(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := models.CreateTestUser()
	assert.NoError(t, err)
	testUserSession, err := models.CreateTestUserSession(user.ID)
	assert.NoError(t, err)

	test.Router.HandleFunc("/logout", postLogout).Methods("POST")
	test.Router.Use(middlewares.SetSession)

	t.Run("Should logout authenticated user", func(t *testing.T) {
		req, err := http.NewRequest("POST", test.Server.URL+"/logout", nil)
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status code to be 200")

		_, err = models.GetSessionBySelector(testUserSession.Selector)
		assert.Equal(t, redis.Nil, err, "Expect session to no longer exist")
	})

	t.Run("Should respond Unauthorized if user is not authenticated", func(t *testing.T) {
		resp, err := http.Post(test.Server.URL+"/logout", "", nil)
		assert.NoError(t, err, "HTTP Post should have no error")

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "Expect status code to be 401")
	})
}
