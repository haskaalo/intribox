package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Album struct {
	ID          uuid.UUID `json:"id" db:"id" pg:",pk,type:uuid default uuid_generate_v4()"`
	OwnerID     int       `json:"ownerid" db:"ownerid"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func InsertNewAlbum(a *Album) (id uuid.UUID, err error) {
	return a.insertNewAlbum(db)
}

func (a *Album) insertNewAlbum(q sqlx.Ext) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO album (ownerid, title, description) VALUES ($1, $2, $3) RETURNING id`
	err := sqlx.Get(q, &id, query, a.OwnerID, a.Title, a.Description)

	return id, knownDatabaseError(err)
}
