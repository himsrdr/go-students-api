package db

import (
	"database/sql"

	"github.com/himsrdr/students-api/internal/config"
	types "github.com/himsrdr/students-api/internal/type"
	_ "github.com/lib/pq"
)

type DB struct {
	DB  *sql.DB
	cfg *config.Config
}

func New(cfg *config.Config) (*DB, error) {
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
		DB:  db,
		cfg: cfg,
	}, nil
}

func (d *DB) GetStudentById(id int64) (types.Student, error) {
	var student types.Student
	query := "select * from students where id = $1 limit 1;"
	err := d.DB.QueryRow(query, id).Scan(&student.Id, &student.Name, &student.Age, &student.Email)
	if err != nil {
		return student, err
	}
	return student, nil
}
func (d *DB) CreateStudent(name, email string, age int) (int64, error) {

	query := "INSERT INTO STUDENTS (Name, age, email) VALUES ($1, $2, $3) RETURNING id"
	var id int64
	err := d.DB.QueryRow(query, name, age, email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}
