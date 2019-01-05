package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/modules/test"
	"github.com/haskaalo/intribox/request"
	"github.com/stretchr/testify/assert"
)

func echoUser(w http.ResponseWriter, r *http.Request) {
	data := request.GetSession(r)
	fmt.Fprintf(w, strconv.Itoa(data.UserID))
}

func TestRequireSession(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	test.Router.HandleFunc("/test/echo/session", echoUser)
	test.Router.Use(SetSession, RequireSession)

	selector, v, err := models.InitiateSession(1)
	assert.NoError(t, err)

	t.Run("Should allow user with a session", func(t *testing.T) {
		req, _ := http.NewRequest("GET", test.Server.URL+"/test/echo/session", nil)
		req.Header.Set(models.SessionHeaderName, selector+"."+v)

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.Nil(t, err)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, strconv.Itoa(1), string(body))
	})

	t.Run("Should not allow user with a session", func(t *testing.T) {
		req, _ := http.NewRequest("GET", test.Server.URL+"/test/echo/session", nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
