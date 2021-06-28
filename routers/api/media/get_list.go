package media

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/storage"
	"github.com/rs/zerolog/log"
)

const defaultMaxLength = 25 // Per page
const maxLengthLimit = 100

type getListResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	UploadedTime int64     `json:"uploaded_time"`
	Size         int64     `json:"size"`
	DownloadURL  string    `json:"download_url"`
}

// List of all media an user has
func getList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	maxLength := r.URL.Query().Get("maxLength") // Number of media url displayed

	// Set default value
	if page == "" {
		page = "1"
	}
	if maxLength == "" {
		maxLength = strconv.Itoa(defaultMaxLength)
	}

	// Convert to int
	numPage, err := strconv.Atoi(page)
	if err != nil {
		response.InvalidParameter(w, "page")
		return
	}

	numMaxLength, err := strconv.Atoi(maxLength)
	if err != nil {
		response.InvalidParameter(w, "maxLength")
		return
	}

	// Check if maxLength isn't out of [0, maxLengthLimit]
	if (numMaxLength > maxLengthLimit) || (numMaxLength <= 0) {
		response.InvalidParameter(w, "maxLength")
		return
	}

	if numPage <= 0 {
		response.InvalidParameter(w, "page")
		return
	}

	userSession := request.GetSession(r)
	mediaList, err := models.GetListMedia(userSession.UserID, numMaxLength, numPage)
	if err == models.ErrRecordNotFound {
		response.Respond(w, &[]getListResponse{}, http.StatusOK)
		return
	} else if err != nil {
		response.InternalError(w)
		return
	}

	formattedResponse := []getListResponse{}
	for _, value := range *mediaList {
		mediaObjectURL, err := storage.Remote.GetReadObjectURL(value.GetMediaPath(), value.ID)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to GetReadObjectURL from remote")
			response.InternalError(w)
			return
		}
		media := getListResponse{
			ID:           value.ID,
			Name:         value.Name,
			UploadedTime: value.UploadedTime.Unix(),
			Size:         value.Size,
			DownloadURL:  mediaObjectURL,
		}
		formattedResponse = append(formattedResponse, media)
	}

	response.Respond(w, &formattedResponse, http.StatusOK)
}
