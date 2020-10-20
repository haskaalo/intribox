package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// Song SQL Table song
type Song struct {
	ID       string    `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	ObjectID string    `json:"objectid" db:"objectid"`
	Ext      string    `json:"ext" db:"ext"`
	OwnerID  int       `json:"ownerid" db:"ownerid"`
	UploadAt time.Time `json:"uploadat" db:"uploadat"`
	FileHash string    `json:"filehash" db:"filehash"`
	Size     int64     `json:"size" db:"size"`
}

// InsertNewSong Insert a new song
func (s *Song) InsertNewSong() (songid int, err error) {
	return s.insertNewSong(db)
}

func (s *Song) insertNewSong(q sqlx.Ext) (int, error) {
	var id int
	query := `INSERT INTO song (name, objectid, ext, ownerid, filehash, size) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := sqlx.Get(q, &id, query, s.Name, s.ObjectID, s.Ext, s.OwnerID, s.FileHash, s.Size)

	return id, knownDatabaseError(err)
}

// SongHashExist Does Song with following hash and ownerid exist
func SongHashExist(ownerid int, hash string) (bool, error) {
	return songHashExist(db, ownerid, hash)
}

func songHashExist(q sqlx.Ext, ownerid int, hash string) (bool, error) {
	var exist bool
	query := `SELECT EXISTS (SELECT 1 FROM song WHERE ownerid=$1 AND filehash=$2)`
	err := sqlx.Get(q, &exist, query, ownerid, hash)

	return exist, knownDatabaseError(err)
}

// GetSongByID Select a song with a ID
func GetSongByID(songid int) (*Song, error) {
	return getSongByID(db, songid)
}

func getSongByID(q sqlx.Ext, songid int) (*Song, error) {
	song := &Song{}
	query := `SELECT * FROM song WHERE id=$1`
	err := sqlx.Get(q, song, query, songid)

	return song, knownDatabaseError(err)
}

// GetSongPath Get song path based on ownerID and objectID
func (s *Song) GetSongPath() string {
	return fmt.Sprintf("%o/song/%s.%s", s.OwnerID, s.ObjectID, s.Ext)
}
