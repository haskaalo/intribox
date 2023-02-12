package album

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
)

type postNewAlbumResponse struct {
	ID uuid.UUID `json:"id"`
}

type postNewResponseParams struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	MediaIDs    []uuid.UUID `json:"media_ids"`
}

func postNewAlbum(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	params := new(postNewResponseParams)
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		response.InvalidParameter(w, "body")
		return
	}

	if params.Title == "" {
		response.InvalidParameter(w, "title")
		return
	}

	userSession := request.GetSession(r)

	// Create an empty album
	album := new(models.Album)
	album.OwnerID = userSession.UserID
	album.Title = params.Title
	album.Description = params.Description

	albumID, err := models.InsertNewAlbum(album)
	if err != nil {
		response.InternalError(w)
		return
	}

	// Insert all pictures into the album

	err = models.InsertNewAlbumMedias(albumID, params.MediaIDs)
	if err != nil {
		response.InternalError(w)
		return
	}

	// Response
	responseVal := postNewAlbumResponse{ID: albumID}
	response.Respond(w, &responseVal, http.StatusOK)
}
