package media

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/haskaalo/intribox/middlewares"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	err = models.DeleteAllMedias()
	assert.NoError(t, err)

	testUser, err := models.CreateTestUser()
	assert.NoError(t, err)

	testUserSession, err := models.CreateTestUserSession(testUser.ID)
	assert.NoError(t, err)

	test.Router.HandleFunc("/list", getList).Methods("GET")
	test.Router.Use(middlewares.SetSession)

	t.Run("Should return invalid parameter if maxLength is out of defined interval", func(t *testing.T) {
		// Prepare request
		req, err := http.NewRequest("GET", test.Server.URL+"/list", nil)
		assert.NoError(t, err)
		q := req.URL.Query()
		q.Add("page", "1")
		q.Add("maxLength", "100000")
		req.URL.RawQuery = q.Encode()
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		// Execute request
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Get should have no error")
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Should return invalid parameter if page is invalid", func(t *testing.T) {
		// Prepare request
		req, err := http.NewRequest("GET", test.Server.URL+"/list", nil)
		assert.NoError(t, err, "Request should have no error")
		q := req.URL.Query()
		q.Add("page", "-1")
		q.Add("maxLength", "25")
		req.URL.RawQuery = q.Encode()
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		// Execute request
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Get should have no error")
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Should successfully return a media list", func(t *testing.T) {
		// Prepare request
		req, err := http.NewRequest("GET", test.Server.URL+"/list", nil)
		assert.NoError(t, err, "Request should have no error")
		q := req.URL.Query()
		q.Add("page", "1")
		q.Add("maxLength", "25")
		req.URL.RawQuery = q.Encode()
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		// Generate Random medias
		allMediaInDatabase := models.GenerateRandomMedia(25, testUser.ID)

		// Execute request
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status code to be 200")

		// Get json
		body, _ := io.ReadAll(resp.Body)
		mediaListInRequest := new([]getListResponse)
		err = json.Unmarshal(body, mediaListInRequest)
		assert.NoError(t, err, "Decoding the JSON should have 0 errors")

		// Check if everything matches descending
		for index, value := range *mediaListInRequest {
			assert.Equal(t, allMediaInDatabase[len(allMediaInDatabase)-index-1].Name, value.Name)
		}
	})
}
