package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// MediaType Custom type for determining if a media file is a picture or video
type MediaType string

// Value return the value of MediaType as byte type
func (m MediaType) Value() (driver.Value, error) {
	return []byte(m), nil
}

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
func (s *Media) InsertNewMedia() (id uuid.UUID, err error) {
	return s.insertNewMedia(db)
}

func (s *Media) insertNewMedia(q sqlx.Ext) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO media (id, name, type, ownerid, filehash, size) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := sqlx.Get(q, &id, query, s.ID, s.Name, s.Type, s.OwnerID, s.FileHash, s.Size)

	return id, knownDatabaseError(err)
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
