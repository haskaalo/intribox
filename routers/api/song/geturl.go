package song

import (
	"encoding/json"
	"net/http"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/storage"
	"github.com/rs/zerolog/log"
)

type getSongURLParams struct {
	SongID int `json:"songid"`
}

func getSongURL(w http.ResponseWriter, r *http.Request) {
	params := new(getSongURLParams)
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		response.InternalError(w) // Probably change depending on the error.
		return
	}

	song, err := models.GetSongByID(params.SongID)
	if err == models.ErrRecordNotFound {
		response.NotFound(w)
		return
	} else if err != nil {
		log.Warn().Err(err).Msg("Error while trying to fetch a specified song from database")
		response.InternalError(w)
		return
	}

	songObjectURL, err := storage.Remote.GetReadObjectURL(song.GetSongPath())
	if err != nil {
		log.Warn().Err(err).Msg("Failed to GetReadObjectURL from remote")
		response.InternalError(w)
		return
	}

	response.Respond(w, &response.M{
		"url": songObjectURL,
	}, http.StatusOK)
}