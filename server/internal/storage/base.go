package storage

import (
	"context"
	"database/sql"
)

// Storage makes CRUD-queries.
type Storage struct {
	db Database
}

func NewStorage() *Storage {
	db := Database{}
	db.Run()
	return &Storage{db: db}
}

func (s *Storage) getValueFromRows(rows *sql.Rows) (result string, err error) {
	defer s.db.closeRows(rows)
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *Storage) receiveRows(
	ctx context.Context, query string, args ...any,
) (*sql.Rows, error) {
	rows, err := s.db.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *Storage) getMeta(value sql.NullString) string {
	if value.Valid {
		return value.String
	} else {
		return ""
	}
}
