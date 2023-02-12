package album

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
	"github.com/rs/zerolog/log"
)

type getAlbumListResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func getAlbumList(w http.ResponseWriter, r *http.Request) {
	userSession := request.GetSession(r)

	albums, err := models.GetAlbumListByOwnerID(userSession.UserID)
	if err != nil && err != models.ErrRecordNotFound {
		response.InternalError(w)
		log.Warn().Err(err).Msg("Failed to contact database on GetAlbumListByOwner")
	}

	responseStruct := []getAlbumListResponse{}

	for _, album := range *albums {
		responseStruct = append(responseStruct, getAlbumListResponse{
			ID:          album.ID,
			Title:       album.Title,
			Description: album.Description,
			CreatedAt:   album.CreatedAt,
		})
	}

	response.Respond(w, &responseStruct, http.StatusOK)
}
