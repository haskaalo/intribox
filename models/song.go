package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Song SQL Table song
type Song struct {
	ID       string    `json:"id" db:"id"`
	OwnerID  int       `json:"ownerid" db:"ownerid"`
	UploadAt time.Time `json:"uploadat" db:"uploadat"`
	FileHash string    `json:"filehash" db:"filehash"`
	FilePath string    `json:"filepath" db:"filepath"`
	Size     int64     `json:"size" db:"size"`
}

// SongDirPrefix used in storage %s represent userID
const SongDirPrefix = "%o/song"

// InsertNewSong Insert a new song
func (s *Song) InsertNewSong() (songid int, err error) {
	return s.insertNewSong(db)
}

func (s *Song) insertNewSong(q sqlx.Ext) (int, error) {
	var id int
	query := `INSERT INTO song (ownerid, filehash, filepath, size) VALUES ($1, $2, $3, $4) RETURNING id`

	err := sqlx.Get(q, &id, query, s.OwnerID, s.FileHash, s.FilePath, s.Size)

	return id, err
}
