package song

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/modules/storage"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
	"github.com/rs/zerolog/log"
)

func postNew(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, config.Server.MaxSongSize)
	defer r.Body.Close()

	session := request.GetSession(r)

	songName := r.Header.Get("X-Song-Name")
	if filepath.Ext(songName) == "" {
		response.InvalidParameter(w, "X-Song-Name")
		return
	}

	if request.RequireContentType("application/octet-stream", r) == false {
		response.InvalidParameter(w, "Content-Type")
		return
	}

	object, err := storage.CurrentRemote.WriteFile(songName, fmt.Sprintf(models.SongDirPrefix, session.UserID), r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error while writing file to remote")
		response.InternalError(w)
		return
	}

	song := &models.Song{
		Name:     fileNameNoExt(songName),
		OwnerID:  session.UserID,
		FileHash: object.SHA256,
		FilePath: object.Path,
		Size:     object.Size,
	}
	songid, err := song.InsertNewSong()
	if err != nil {
		log.Error().Err(err).Msg("Error while inserting song info to database")
		response.InternalError(w)

		err = storage.CurrentRemote.RemoveFile(song.FilePath)
		if err != nil {
			log.Error().Err(err).Msg("Error while removing song info to database after having error with database")
		}
		return
	}

	response.Respond(w, &response.M{
		"id": songid,
	}, 200)
}

func fileNameNoExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
