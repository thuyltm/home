package service

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
)

type CockroachDBMessageService struct {
	db *sql.DB
}

func NewCockroachDBMessageService(db *sql.DB) *CockroachDBMessageService {
	return &CockroachDBMessageService{db: db}
}

func (s *CockroachDBMessageService) Close() {
	s.db.Close()
}

func (s *CockroachDBMessageService) CountMessages() (int, error) {
	rows, err := s.db.Query("SELECT COUNT(*) FROM message")
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
		rows.Close()
	}
	return count, nil
}

func (s *CockroachDBMessageService) CreateMessage(value string) error {
	err := crdb.ExecuteTx(context.Background(), s.db, nil, func(tx *sql.Tx) error {
		_, err := tx.Exec(
			"Insert into message (value) values ($1) on conflict (value) do update set value = excluded.value",
			value,
		)
		return err
	})
	return err
}
