package media

import (
	"encoding/json"
	"net/http"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/storage"
	"github.com/rs/zerolog/log"
)

type getMediaURLParams struct {
	MediaID int `json:"mediaid"`
}

type getMediaURLResponse struct {
	URL string `json:"url"`
}

func getMediaURL(w http.ResponseWriter, r *http.Request) {
	userSession := request.GetSession(r)
	params := new(getMediaURLParams)
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		response.InternalError(w) // TODO: Change the error depending on encoding/json error
		return
	}

	media, err := models.GetMediaByID(params.MediaID, userSession.UserID)
	if err == models.ErrRecordNotFound {
		response.NotFound(w)
		return
	} else if err != nil {
		log.Warn().Err(err).Msg("Error while trying to fetch a specified media from database")
		response.InternalError(w)
		return
	}

	mediaObjectURL, err := storage.Remote.GetReadObjectURL(media.GetMediaPath())
	if err != nil {
		log.Warn().Err(err).Msg("Failed to GetReadObjectURL from remote")
		response.InternalError(w)
		return
	}

	response.Respond(w, &getMediaURLResponse{
		URL: mediaObjectURL,
	}, http.StatusOK)
}
