package sqlite

import (
	"database/sql"

	"github.com/himsrdr/students-api/internal/config"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

func NewSqlite(cfg *config.Config) (*DB, error) {
	db, err := sql.Open("postgres", cfg.PostgresUrl)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS STUDENTS (
    id SERIAL PRIMARY KEY,  -- Use SERIAL for auto-incrementing IDs
    Name TEXT,
    age INTEGER,
    email TEXT
);
`)
	if err != nil {
		return nil, err
	}
	return &DB{
		DB: db,
	}, nil

}
