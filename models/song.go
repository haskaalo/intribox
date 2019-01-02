package models

import "time"

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
func (u *Song) InsertNewSong() error {
	return u.insertNewSong(db)
}

func (u *Song) insertNewSong(q Querier) error {
	_, err := q.NamedExec("INSERT INTO song (ownerid, filehash, filepath, size) VALUES (:ownerid, :filehash, :filepath, :size)",
		map[string]interface{}{
			"ownerid":  u.OwnerID,
			"filehash": u.FileHash,
			"filepath": u.FilePath,
			"size":     u.Size,
		})
	return err
}
