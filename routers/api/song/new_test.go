package song

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/haskaalo/intribox/middlewares"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/modules/test"
	"github.com/haskaalo/intribox/response"
	"github.com/stretchr/testify/assert"
)

func TestPostNew(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := test.CreateTestUser()
	assert.NoError(t, err)

	testUserSession, err := test.CreateTestUserSession(user.ID)
	assert.NoError(t, err)

	test.Router.HandleFunc("/new", postNew).Methods("POST")
	test.Router.Use(middlewares.SetSession)

	t.Run("Should upload new song with no error", func(t *testing.T) {
		req, err := http.NewRequest("POST", test.Server.URL+"/new", bytes.NewBuffer([]byte("Some sort of content")))
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)
		req.Header.Add("Content-Type", SongContentType)
		req.Header.Add(SongNameHeaderName, "Testing - A song.mp3")

		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status code to be 200")

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err, "Should not have an error while reading request body")

		reqBody := response.M{}
		err = json.Unmarshal(body, &reqBody)
		assert.NoError(t, err)

		song, err := models.GetSongByID(int(reqBody["id"].(float64)), user.ID)
		assert.NoError(t, err)

		assert.Equal(t, "Testing - A song.mp3", song.Name+"."+song.Ext, "The song should exist in the database")
		// TODO: Check if it exist in storage
	})

	t.Run("Should return invalid parameter if no song name is in the header", func(t *testing.T) {
		req, err := http.NewRequest("POST", test.Server.URL+"/new", bytes.NewBuffer([]byte("Some sort of content")))
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expect status code to be 400")
	})

	t.Run("Should return invalid parameter if content-type doesn't match SongContentType", func(t *testing.T) {
		selector, validator, err := models.InitiateSession(user.ID)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", test.Server.URL+"/new", bytes.NewBuffer([]byte("Some sort of content")))
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, selector+"."+validator)
		req.Header.Add(SongNameHeaderName, "Testing - A song.mp3")
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expect status code to be 400")
	})
}
