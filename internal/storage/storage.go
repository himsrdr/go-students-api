package storage

import types "github.com/himsrdr/students-api/internal/type"

type Storage interface {
	CreateStudent(name, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
}
