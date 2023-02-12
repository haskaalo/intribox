package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AlbumMedia struct {
	ID      uuid.UUID `json:"id" db:"id" pg:",pk,type:uuid default uuid_generate_v4()"`
	AlbumID uuid.UUID `json:"albumid" db:"albumid" pg:",type:uuid"`
	MediaID uuid.UUID `json:"mediaid" db:"mediaid" pg:",type:uuid"`
}

func InsertNewAlbumMedias(albumID uuid.UUID, mediaIDs []uuid.UUID) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, mediaID := range mediaIDs {
		_, err = tx.Exec("INSERT INTO album_media (albumid, mediaid) VALUES ($1, $2)", albumID, mediaID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}

func GetAlbumMediasByAlbumID(albumID uuid.UUID) ([]AlbumMedia, error) {
	return getAlbumMediasByAlbumID(db, albumID)
}

func getAlbumMediasByAlbumID(q sqlx.Ext, albumID uuid.UUID) ([]AlbumMedia, error) {
	list := new([]AlbumMedia)
	query := `SELECT * FROM album_media
	          WHERE albumid=$1`

	err := sqlx.Select(q, list, query, albumID)

	return *list, knownDatabaseError(err)
}
