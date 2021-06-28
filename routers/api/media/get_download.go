package media

import (
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/storage"
	"github.com/rs/zerolog/log"
)

// This API Call is only used if remote is local
func getDownload(w http.ResponseWriter, r *http.Request) {
	if config.Storage.RemoteName != "local" {
		response.NotImplemented(w)
		return
	}
	mediaIDStr := r.URL.Query().Get("id")
	if mediaIDStr == "" {
		response.InvalidParameter(w, "id")
		return
	}

	mediaID, err := uuid.Parse(mediaIDStr)
	if err != nil {
		response.InvalidParameter(w, "id")
		return
	}

	userSession := request.GetSession(r)

	media, err := models.GetMediaByID(mediaID, userSession.UserID)
	if err == models.ErrRecordNotFound {
		response.NotFound(w)
		return
	} else if err != nil {
		log.Warn().Err(err).Msg("Error while trying to fetch a specified media from database")
		response.InternalError(w)
		return
	}

	reader, err := storage.Remote.ReadObject(media.GetMediaPath())
	if err != nil {
		log.Warn().Err(err).Msg("Error while trying to fetch a specified media from local remote")
		response.InternalError(w)
		return
	}

	w.Header().Set("Content-Type", media.Type)
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, reader)
	if err != nil {
		response.InternalError(w)
		return
	}
}
