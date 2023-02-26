package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/haskaalo/intribox/utils"
	"github.com/jmoiron/sqlx"
)

// Media SQL Table
type Media struct {
	ID           uuid.UUID `json:"id" db:"id" pg:",pk,type:uuid default uuid_generate_v4()"`
	Name         string    `json:"name" db:"name"`
	Type         string    `json:"type" db:"type"`
	OwnerID      int       `json:"ownerid" db:"ownerid"`
	UploadedTime time.Time `json:"uploaded_time" db:"uploaded_time"`
	FileHash     string    `json:"filehash" db:"filehash"`
	Size         int64     `json:"size" db:"size"`
}

// InsertNewMedia Insert a new media file
func (s *Media) InsertNewMedia() (err error) {
	return s.insertNewMedia(db)
}

func (s *Media) insertNewMedia(q sqlx.Ext) error {
	query := `INSERT INTO media (id, name, type, ownerid, filehash, size) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	err := sqlx.Get(q, s, query, s.ID, s.Name, s.Type, s.OwnerID, s.FileHash, s.Size)

	return knownDatabaseError(err)
}

// MediaHashExist Does a media file with following hash and ownerid exist
func MediaHashExist(ownerid int, hash string) (bool, error) {
	return mediaHashExist(db, ownerid, hash)
}

func mediaHashExist(q sqlx.Ext, ownerid int, hash string) (bool, error) {
	var exist bool
	query := `SELECT EXISTS (SELECT 1 FROM media WHERE ownerid=$1 AND filehash=$2)`
	err := sqlx.Get(q, &exist, query, ownerid, hash)

	return exist, knownDatabaseError(err)
}

// GetMediaByID Select a media file with a ID
func GetMediaByID(mediaid uuid.UUID, ownerid int) (*Media, error) {
	return getMediaByID(db, mediaid, ownerid)
}

func getMediaByID(q sqlx.Ext, id uuid.UUID, ownerid int) (*Media, error) {
	media := new(Media)
	query := `SELECT * FROM media WHERE id=$1 AND ownerid=$2`
	err := sqlx.Get(q, media, query, id, ownerid)

	return media, knownDatabaseError(err)
}

// GetMediaPath Get media file path based on ownerID and ID
func (s *Media) GetMediaPath() string {
	return fmt.Sprintf("%o/media/%s", s.OwnerID, s.ID.String())
}

// DeleteAllMedias Only should be used for testing (TODO: Obviously doesn't belong here)
func DeleteAllMedias() error {
	_, err := db.NamedExec("DELETE FROM media", map[string]interface{}{})
	return err
}

func getListMedia(q sqlx.Ext, ownerid int, maxLength int, page int) (*[]Media, error) {
	mediaList := new([]Media)
	query := `SELECT * FROM media 
			WHERE ownerid=$1 
			ORDER BY uploaded_time DESC
			OFFSET ($2 * ($3 - 1)) LIMIT $2`
	err := sqlx.Select(q, mediaList, query, ownerid, maxLength, page)

	return mediaList, knownDatabaseError(err)
}

func GetListMedia(ownerid int, maxLength int, page int) (*[]Media, error) {
	return getListMedia(db, ownerid, maxLength, page)
}

// GenerateRandomMedia Obviously for testing purposes
func GenerateRandomMedia(n int, ownerID int) []Media {
	allMediaInDatabase := []Media{}
	// Insert 25 random
	for i := 0; i < n; i++ {
		mediaTest := &Media{
			ID:           uuid.New(),
			Name:         utils.RandString(5),
			Type:         "image/png",
			OwnerID:      ownerID,
			UploadedTime: time.Now().Add(time.Duration(5*i) * time.Minute),
			FileHash:     utils.SHA256([]byte(utils.RandString(5))),
			Size:         420,
		}
		allMediaInDatabase = append(allMediaInDatabase, *mediaTest)
		err := mediaTest.InsertNewMedia()
		if err != nil {
			panic(err)
		}
	}

	return allMediaInDatabase
}
