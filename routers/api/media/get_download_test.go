package media

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/middlewares"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/storage"
	"github.com/haskaalo/intribox/storage/backend/local"
	"github.com/haskaalo/intribox/utils/test"
	"github.com/stretchr/testify/assert"
)

func testPostTestingImage(url string, testSession *models.TestingUserSession, t *testing.T) int {
	// Create custom multipart
	reqBody, contentType, _ := createNewTestMultipart("testimage.png", "image/png")

	// Prepare request
	req, err := http.NewRequest("POST", url, reqBody)
	assert.NoError(t, err, "Request should have no error")
	req.Header.Add(models.SessionHeaderName, testSession.FullSessionToken)
	req.Header.Add("Content-Type", contentType)

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)

	assert.NoError(t, err, "HTTP Post should have no error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status code to be 200")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "Should not have an error while reading request body")

	resBody := &postNewResponse{}
	err = json.Unmarshal(body, &resBody)
	assert.NoError(t, err)

	return resBody.ID
}

func TestGetDownload(t *testing.T) {
	// Because this request only works on remote local
	oldRemoteName := config.Storage.RemoteName
	oldRemote := storage.Remote

	config.Storage.RemoteName = "local"
	storage.Remote = new(local.R)
	defer func() {
		config.Storage.RemoteName = oldRemoteName
		storage.Remote = oldRemote
	}()

	// Setup server and test user
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := models.CreateTestUser()
	assert.NoError(t, err)

	testSession, err := models.CreateTestUserSession(user.ID)
	assert.NoError(t, err)

	// Setup router
	test.Router.HandleFunc("/download", getDownload).Methods("GET")
	test.Router.HandleFunc("/new", postNew).Methods("POST")
	test.Router.Use(middlewares.SetSession)
	downloadTestingURL := test.Server.URL + "/download"
	newTestingURL := test.Server.URL + "/new"

	t.Run("Should successfully return a picture or video", func(t *testing.T) {
		// Create new testing image
		mediaID := testPostTestingImage(newTestingURL, testSession, t)

		// Prepare request
		req, err := http.NewRequest("GET", downloadTestingURL+"?mediaid="+strconv.Itoa(mediaID), nil)
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testSession.FullSessionToken)

		// Do request
		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err, "HTTP Get should have no error")

		assert.Equal(t, http.StatusOK, resp.StatusCode, "Status should be 200")
	})

	t.Run("Should return 404 for an unknown picture/video", func(t *testing.T) {
		// Prepare request
		req, err := http.NewRequest("GET", downloadTestingURL+"?mediaid=78236478", nil)
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testSession.FullSessionToken)

		// Do request
		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err, "HTTP Get should have no error")

		assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Status should be 404")
	})
}
