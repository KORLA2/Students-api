package storage

import "myapp/student/Type"

type Storage interface {
	CreateStudent(name string, age int, email string) (int64, error)
	GetStudent(id int64) (Type.Student, error)
	ListAllStudents() ([]Type.Student, error)
	UpdateStudent(id int64, name string, age int, email string) (int64, error)
	DeleteStudent(id int64) (int64, error)
}
