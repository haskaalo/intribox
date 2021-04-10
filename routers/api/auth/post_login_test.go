package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := models.CreateTestUser()
	assert.NoError(t, err)

	test.Router.HandleFunc("/login", postLogin).Methods("POST")

	t.Run("Should authenticate user with correct password", func(t *testing.T) {
		jsonLoginParams, err := json.Marshal(&loginParams{
			Email:    user.Email,
			Password: models.TestUserPassword,
		})
		assert.NoError(t, err)

		resp, err := http.Post(test.Server.URL+"/login", "application/json", bytes.NewReader(jsonLoginParams))
		assert.NoError(t, err, "HTTP post should not have error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status to be OK")

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err, "Should not have an error while reading request body")

		reqBody := response.M{}
		err = json.Unmarshal(body, &reqBody)
		assert.NoError(t, err)

		assert.NotEmpty(t, reqBody["apiToken"], "Response JSON body should contain a apiToken field not empty")
	})

	t.Run("Should not authenticate user with incorrect password", func(t *testing.T) {
		jsonLoginParams, err := json.Marshal(&loginParams{
			Email:    user.Email,
			Password: "obviously not the correct password",
		})
		assert.NoError(t, err)

		resp, err := http.Post(test.Server.URL+"/login", "application/json", bytes.NewReader(jsonLoginParams))
		assert.NoError(t, err, "HTTP post should not have error")

		assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Expect status to be Not Found")

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err, "Should not have an error while reading request body")

		reqBody := response.M{}
		err = json.Unmarshal(body, &reqBody)
		assert.NoError(t, err)

		assert.Nil(t, reqBody["apiToken"], "Body should not contains a api token")
	})
}
