package album

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

func TestGetAlbumList(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := models.CreateTestUser()
	assert.NoError(t, err)

	testUserSession, err := models.CreateTestUserSession(user.ID)
	assert.NoError(t, err)

	test.Router.HandleFunc("/albumlist", getAlbumList).Methods("GET")
	test.Router.Use(middlewares.SetSession)

	t.Run("Should respond with all user album", func(t *testing.T) {
		// Create 2 testing empty album
		album := new(models.Album)
		album.OwnerID = user.ID
		album.Title = "Test album 1"
		album.Description = "Some description"

		albumID_1, err := models.InsertNewAlbum(album)
		assert.NoError(t, err)

		album.Title = "Test album 2"
		albumID_2, err := models.InsertNewAlbum(album)
		assert.NoError(t, err)

		// Prepare request
		req, err := http.NewRequest("GET", test.Server.URL+"/albumlist", nil)
		assert.NoError(t, err)

		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)

		// Execute test request
		client := &http.Client{}
		resp, err := client.Do(req)

		// Testing done on response
		assert.NoError(t, err, "HTTP get should not have error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status to be OK")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err, "Should not have an error while reading request body")

		resBody := []getAlbumListResponse{}
		err = json.Unmarshal(body, &resBody)
		assert.NoError(t, err)

		seenAlbum := 0
		for _, album := range resBody {
			if album.ID == albumID_1 {
				assert.Equal(t, "Test album 1", album.Title)
				seenAlbum += 1
			} else if album.ID == albumID_2 {
				assert.Equal(t, "Test album 2", album.Title)
				seenAlbum += 1
			}
		}

		assert.Equal(t, 2, seenAlbum, "Should have 2 albums in the response")
	})
}
