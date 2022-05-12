package model

import (
	"database/sql"
)

type Thumbnail struct {
	IDVideo        sql.NullString `db:"id"`
	ThumbnailImage []byte         `db:"thumbnail"`
	CreatedAt      sql.NullTime   `db:"created_at"`
}
