package song

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/modules/context"
	"github.com/haskaalo/intribox/modules/storage"
	"github.com/haskaalo/intribox/response"
	"github.com/rs/zerolog/log"
)

func postNew(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, config.Server.MaxSongSize)
	defer r.Body.Close()

	session := context.GetSession(r)

	songName := r.Header.Get("X-Song-Name")
	if filepath.Ext(songName) == "" {
		response.InvalidParameter(w, "X-Song-Name")
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/octet-stream" {
		response.InvalidParameter(w, "Content-Type")
		return
	}

	object, err := storage.CurrentRemote.WriteFile(songName, fmt.Sprintf(models.SongDirPrefix, session.UserID), r.Body)
	if err != nil {
		response.InternalError(w)
		log.Error().Err(err).Msg("Error while writing file to remote")
		return
	}

	song := &models.Song{
		OwnerID:  session.UserID,
		FileHash: object.SHA256,
		FilePath: object.Path,
		Size:     object.Size,
	}

	err = song.InsertNewSong()
	if err != nil {
		response.InternalError(w)
		log.Error().Err(err).Msg("Error while inserting song info to database")
		return
		// storage.CurrentRemote.RemoveFile(...)
	}

	response.Respond(w, &response.M{
		"ok": 1,
	}, 200) // replace with song id and more
}
