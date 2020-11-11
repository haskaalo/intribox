package song

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/middlewares"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/modules/test"
	"github.com/stretchr/testify/assert"
)

func TestGetSongURL(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := models.CreateTestUser()
	assert.NoError(t, err)

	userSession, err := models.CreateTestUserSession(user.ID)
	assert.NoError(t, err)

	test.Router.HandleFunc("/get_songurl", getSongURL).Methods("GET")
	test.Router.Use(middlewares.SetSession)
	testingURL := test.Server.URL + "/get_songurl"

	t.Run("Should return a valid song download url", func(t *testing.T) {
		// Inserting a fake song in database
		fakeSong := &models.Song{
			Name:     "Darude - Sandstorm",
			ObjectID: uuid.New().String(),
			Ext:      "mp3",
			OwnerID:  user.ID,
			FileHash: "ab43487f946e97f24100685cb1d167024eb9dce910c18686feecf814bccc1749",
			Size:     420,
		}

		fakeSongID, err := fakeSong.InsertNewSong()
		assert.NoError(t, err)
		jsonGetURLParams, err := json.Marshal(&getSongURLParams{
			SongID: fakeSongID,
		})
		assert.NoError(t, err)

		// Actual test
		httpRequest, err := http.NewRequest("GET", testingURL, bytes.NewReader(jsonGetURLParams))
		assert.NoError(t, err, "Request should have no error when created")
		httpRequest.Header.Add(models.SessionHeaderName, userSession.FullSessionToken)

		httpClient := &http.Client{}
		resp, err := httpClient.Do(httpRequest)
		assert.NoError(t, err, "HTTP request should have no error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status code to be 200 (OK)")

		defer resp.Body.Close()
		reqBody := new(getSongURLResponse)
		err = json.NewDecoder(resp.Body).Decode(reqBody)
		assert.NoError(t, err)
		assert.NotEmpty(t, reqBody.URL, "Expect the Song URL string to not be empty")

		// TODO: assert if the returned URL work
	})

	t.Run("Should return 404 if the song doesn't exist", func(t *testing.T) {
		jsonGetURLParams, err := json.Marshal(&getSongURLParams{
			SongID: 420000, // This song ID doesn't exist in testing database
		})

		httpRequest, err := http.NewRequest("GET", testingURL, bytes.NewReader(jsonGetURLParams))
		assert.NoError(t, err, "Request should have no error when created")
		httpRequest.Header.Add(models.SessionHeaderName, userSession.FullSessionToken)

		httpClient := &http.Client{}
		resp, err := httpClient.Do(httpRequest)

		assert.NoError(t, err, "HTTP request should have no error")
		assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Expect status code to equal 404 (NOT FOUND)")

	})
}
