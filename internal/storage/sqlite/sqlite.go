package sqlite

import (
	"database/sql"

	"github.com/himsrdr/students-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

func NewSqlite(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS STUDENTS(id INTEGER Primary key AutoIncreament ,
	 Name Text ,
	  age integer ,
	   email text
	)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{
		DB: db,
	}, nil

}
