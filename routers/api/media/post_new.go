package media

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/storage"
	"github.com/rs/zerolog/log"
)

type postNewResponse struct {
	ID int `json:"id"`
}

func validNewFileContentType(contentType string) bool {
	s := strings.Split(contentType, "/")
	if len(s) == 0 {
		return false
	} else if len(s) != 2 {
		return false
	}

	if s[0] == "image" || s[0] == "video" {
		return true
	}

	return false
}

func postNew(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, config.Server.MaxMediaSize)
	session := request.GetSession(r)

	file, handler, err := r.FormFile("file")
	if err != nil {
		response.InvalidParameter(w, "body")
		return
	}
	defer file.Close()

	contentType := r.FormValue("content-type")
	if !validNewFileContentType(contentType) {
		response.InvalidParameter(w, "filename")
		return
	}

	media := &models.Media{
		Name:     handler.Filename,
		ObjectID: uuid.New().String(),
		Type:     contentType,
		OwnerID:  session.UserID,
	}

	objectWriter, err := storage.Remote.WriteObject(file, media.GetMediaPath())
	if err != nil {
		log.Warn().Err(err).Msg("Error while creating object writer")
		response.InternalError(w)
		return
	}

	exist, err := models.MediaHashExist(session.UserID, objectWriter.SHA256())
	if err != nil {
		log.Warn().Err(err).Msg("Error while querying database")
		if objectWriter.Delete() != nil {
			log.Fatal().Err(err).Msg("Couldn't delete object")
		}
		response.InternalError(w)
		return
	} else if exist {
		if objectWriter.Delete() != nil {
			log.Fatal().Err(err).Msg("Couldn't delete object")
		}
		response.Conflict(w)
		return
	}

	media.FileHash = objectWriter.SHA256()
	media.Size = objectWriter.Size()

	mediaid, err := media.InsertNewMedia()
	if err != nil {
		response.InternalError(w)

		err = objectWriter.Delete()
		if err != nil {
			log.Error().Err(err).Str("path", media.GetMediaPath()).Msg("Cannot remove file from remote after error from writing to database")
		}

		log.Warn().Err(err).Msg("Error while inserting media metadata to database")
		return
	}

	response.Respond(w, &postNewResponse{
		ID: mediaid,
	}, 200)
}
