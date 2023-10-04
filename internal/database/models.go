package database

import (
	"database/sql"
	"time"
)

type ApodDatum struct {
	ID          int32
	Date        time.Time
	Title       string
	Description sql.NullString
	Url         string
	MediaType   string
}
