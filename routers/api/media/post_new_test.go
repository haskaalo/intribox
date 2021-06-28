package media

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/middlewares"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/utils"
	"github.com/haskaalo/intribox/utils/test"
	"github.com/stretchr/testify/assert"
)

// createNewTestMultipart This creates a multipart/form-data to work with in tests
func createNewTestMultipart(fileName string, contentType string) (*bytes.Buffer, string, string) {
	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	part, _ := writer.CreateFormFile("file", fileName)

	body := utils.RandString(500) // Pretend this is the binary data in the picture file
	_, _ = part.Write([]byte(body))

	_ = writer.WriteField("content-type", contentType)
	writer.Close()

	sha256 := utils.SHA256([]byte(body))
	return reqBody, writer.FormDataContentType(), sha256
}

func TestPostNew(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	err := models.DeleteAllUsers()
	assert.NoError(t, err)

	user, err := models.CreateTestUser()
	assert.NoError(t, err)

	testUserSession, err := models.CreateTestUserSession(user.ID)
	assert.NoError(t, err)

	test.Router.HandleFunc("/new", postNew).Methods("POST")
	test.Router.Use(middlewares.SetSession)

	t.Run("Should upload new media with no error", func(t *testing.T) {
		// Create custom multipart
		reqBody, contentType, _ := createNewTestMultipart("testimage.png", "image/png")

		// Prepare request
		req, err := http.NewRequest("POST", test.Server.URL+"/new", reqBody)
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)
		req.Header.Add("Content-Type", contentType)

		// Execute request
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expect status code to be 200")

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err, "Should not have an error while reading request body")

		resBody := response.M{}
		err = json.Unmarshal(body, &resBody)
		assert.NoError(t, err)

		bodyID, err := uuid.Parse(resBody["id"].(string))
		assert.NoError(t, err, "ID must have no error when parsing into uuid type")

		media, err := models.GetMediaByID(bodyID, user.ID)
		assert.NoError(t, err)

		assert.Equal(t, "testimage.png", media.Name, "The picture should exist in the database")
		// TODO: Check if it exist in storage
	})

	t.Run("Should return invalid parameter if no media name is in the body", func(t *testing.T) {
		// Create custom multipart
		reqBody, contentType, _ := createNewTestMultipart("", "video/mp4")

		req, err := http.NewRequest("POST", test.Server.URL+"/new", reqBody)
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)
		req.Header.Add("Content-Type", contentType)
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expect status code to be 400")
	})

	t.Run("Should return invalid parameter if content-type doesn't match multipart/form-data", func(t *testing.T) {
		reqBody, _, _ := createNewTestMultipart("", "video/mp4")

		req, err := http.NewRequest("POST", test.Server.URL+"/new", reqBody)
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err, "HTTP Post should have no error")
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expect status code to be 400")
	})

	t.Run("Should return invalid parameter if content-type field isn't an image or video", func(t *testing.T) {
		reqBody, multipartType, _ := createNewTestMultipart("", "video/mp4")

		req, err := http.NewRequest("POST", test.Server.URL+"/new", reqBody)
		assert.NoError(t, err, "Request should have no error")
		req.Header.Add(models.SessionHeaderName, testUserSession.FullSessionToken)
		req.Header.Add("Content-Type", multipartType)

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expect status code to be 400")
	})
}
