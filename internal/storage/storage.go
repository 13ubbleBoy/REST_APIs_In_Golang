package storage

import "github.com/13ubbleBoy/REST_APIs_In_Golang/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error) // signature of the method to add a student into the database
	GetStudentById(id int64) (types.Student, error)                  // signature of the method to get a student by id
	GetStudents() ([]types.Student, error)                           // signature of the method to get list of all the students
}
