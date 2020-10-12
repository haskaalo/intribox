package song

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/storage"
	"github.com/rs/zerolog/log"
)

// SongNameHeadName The header required to know song name
const SongNameHeaderName = "X-Song-Name"

// SongContentType Content-Type required to post a new song
const SongContentType = "application/octet-stream"

func postNew(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, config.Server.MaxSongSize)
	defer r.Body.Close()

	session := request.GetSession(r)

	songName := r.Header.Get(SongNameHeaderName)
	if filepath.Ext(songName) == "" {
		response.InvalidParameter(w, SongNameHeaderName)
		return
	}

	if request.RequireContentType(SongContentType, r) == false {
		response.InvalidParameter(w, "Content-Type")
		return
	}

	song := &models.Song{
		Name:     fileNameNoExt(songName),
		ObjectID: uuid.New().String(),
		Ext:      filepath.Ext(songName)[1:],
		OwnerID:  session.UserID,
	}

	objectWriter, err := storage.Remote.WriteObject(r.Body, song.GetSongPath())
	if err != nil {
		log.Warn().Err(err).Msg("Error while creating object writer")
		response.InternalError(w)
		return
	}

	exist, err := models.SongHashExist(session.UserID, song.FileHash)
	if err != nil {
		log.Warn().Err(err).Msg("Error while querying database")
		objectWriter.Delete()
		response.InternalError(w)
		return
	} else if exist == true {
		objectWriter.Delete()
		response.Conflict(w)
		return
	}

	song.FileHash = objectWriter.SHA256()
	song.Size = objectWriter.Size()

	songid, err := song.InsertNewSong()
	if err != nil {
		response.InternalError(w)

		err = objectWriter.Delete()
		if err != nil {
			log.Error().Err(err).Str("path", song.GetSongPath()).Msg("Cannot remove file from remote after error from writing to database")
		}

		log.Warn().Err(err).Msg("Error while inserting song metadata to database")
		return
	}

	response.Respond(w, &response.M{
		"id": songid,
	}, 200)
}

func fileNameNoExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
