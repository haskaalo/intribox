package middlewares

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/utils"
	"github.com/haskaalo/intribox/utils/test"
	"github.com/stretchr/testify/assert"
)

func echoSessionData(w http.ResponseWriter, r *http.Request) {
	data := request.GetSession(r)
	fmt.Fprintf(w, data.Selector+"|"+data.Validator+"|"+strconv.Itoa(data.UserID))
}

func TestSetSession(t *testing.T) {
	test.MockServerSetup()
	defer test.MockServerTearDown()

	test.Router.HandleFunc("/test/echo/session", echoSessionData)
	test.Router.Use(SetSession)

	req, _ := http.NewRequest("GET", test.Server.URL+"/test/echo/session", nil)
	selector, v, err := models.InitiateSession(1)
	assert.NoError(t, err)
	req.Header.Set(models.SessionHeaderName, selector+"."+v)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.Nil(t, err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, selector+"|"+utils.SHA1([]byte(v))+"|"+"1", string(body))
}
