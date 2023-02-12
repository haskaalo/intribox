package media

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/middlewares"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestPostNewAlbum(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := models.CreateTestUser()
	assert.NoError(t, err)

	testUserSession, err := models.CreateTestUserSession(user.ID)
	assert.NoError(t, err)

	test.Router.HandleFunc("/newalbum", postNewAlbum).Methods("POST")
	test.Router.Use(middlewares.SetSession)

	t.Run("Should upload new album with no error", func(t *testing.T) {
		// Generate random media for user
		mediaIDs := []uuid.UUID{}
		allMediaInDatabase := models.GenerateRandomMedia(25, user.ID)

		for _, media := range allMediaInDatabase {
			mediaIDs = append(mediaIDs, media.ID)
		}

		// Prepare test request
		requestBody, err := json.Marshal(&postNewResponseParams{
			Title:       "Best of Shamwow",
			Description: "",
			MediaIDs:    mediaIDs,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", test.Server.URL+"/newalbum", bytes.NewReader(requestBody))
		assert.NoError(t, err)

		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		// Execute test request
		client := &http.Client{}
		resp, err := client.Do(req)

		// Testing done on response
		assert.NoError(t, err, "HTTP post should not have error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status to be OK")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err, "Should not have an error while reading request body")

		reqBody := postNewAlbumResponse{}
		err = json.Unmarshal(body, &reqBody)
		assert.NoError(t, err)

		assert.NotEmpty(t, reqBody.ID, "Response body ID should not be null/empty")

		// Check if every photo has been inserted into album
		albumMedias, err := models.GetAlbumMediasByAlbumID(reqBody.ID)
		assert.NoError(t, err)

		initialMediaID := map[string]bool{}
		for _, mediaID := range mediaIDs {
			initialMediaID[mediaID.String()] = false
		}

		for _, albumMedia := range albumMedias {
			val, ok := initialMediaID[albumMedia.MediaID.String()]
			assert.True(t, ok, "A media should've been inserted into album")
			assert.False(t, val, "No duplicate should exist in that album")

			initialMediaID[albumMedia.MediaID.String()] = true
		}

	})

	t.Run("Should not be able to upload album with empty title", func(t *testing.T) {
		// Generate random media for user
		mediaIDs := []uuid.UUID{}

		// Prepare test request
		requestBody, err := json.Marshal(&postNewResponseParams{
			Title:       "",
			Description: "",
			MediaIDs:    mediaIDs,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", test.Server.URL+"/newalbum", bytes.NewReader(requestBody))
		assert.NoError(t, err)

		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		// Execute test request
		client := &http.Client{}
		resp, err := client.Do(req)

		// Testing done on response
		assert.NoError(t, err, "HTTP post should not have error")
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expect status to be 400")
	})
}
