package database

import (
	"context"
	"database/sql"
	"time"
)

const deleteByDate = `

DELETE FROM apod_data WHERE "date" = $1
`

func (q *Queries) DeleteByDate(ctx context.Context, date time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteByDate, date)
	return err
}

const getByDate = `

SELECT id, date, title, description, url, media_type FROM apod_data WHERE date = $1
`

func (q *Queries) GetByDate(ctx context.Context, date time.Time) (ApodDatum, error) {
	row := q.db.QueryRowContext(ctx, getByDate, date)
	var i ApodDatum
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.MediaType,
	)
	return i, err
}

const getByDateRange = `

SELECT id, date, title, description, url, media_type FROM apod_data WHERE date >= $1 AND "date" <= $2
`

type GetByDateRangeParams struct {
	Date   time.Time
	Date_2 time.Time
}

func (q *Queries) GetByDateRange(ctx context.Context, arg GetByDateRangeParams) ([]ApodDatum, error) {
	rows, err := q.db.QueryContext(ctx, getByDateRange, arg.Date, arg.Date_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApodDatum
	for rows.Next() {
		var i ApodDatum
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Title,
			&i.Description,
			&i.Url,
			&i.MediaType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOne = `INSERT INTO apod_data
(id, date, title, description, url, media_type)
VALUES($1, $2, $3, $4, $5, $6)
`

type InsertOneParams struct {
	ID          int32
	Date        time.Time
	Title       string
	Description sql.NullString
	Url         string
	MediaType   string
}

func (q *Queries) InsertOne(ctx context.Context, arg InsertOneParams) error {
	_, err := q.db.ExecContext(ctx, insertOne,
		arg.ID,
		arg.Date,
		arg.Title,
		arg.Description,
		arg.Url,
		arg.MediaType,
	)
	return err
}
