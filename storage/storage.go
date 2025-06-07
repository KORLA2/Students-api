package storage

import "myapp/student/Type"

type Storage interface {
	CreateStudent(name string, age int, email string) (int64, error)
	GetStudent(id int64) (Type.Student, error)
	ListAllStudents() ([]Type.Student, error)
}
