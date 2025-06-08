package sqlite

import (
	"database/sql"
	"fmt"
	"myapp/config"

	"myapp/student/Type"

	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	fmt.Println("COnnncetion to the database", cfg.StoragePath)
	db, err := sql.Open("sqlite", cfg.StoragePath)

	if err != nil {
		fmt.Println("Connection to the database failed")
		return nil, err
	}
	Result, err := db.Exec(`Create Table If Not Exists Student(
	Id Integer Primary Key AutoIncrement,
    Name Text,
	Age Integer,
	Email Text 
	)`)

	if err != nil {

		fmt.Println("Failed to execute create query")
		return nil, err
	}
	fmt.Println("This sis the conncetion result", Result)
	return &Sqlite{
		db,
	}, nil

}

func (s *Sqlite) CreateStudent(name string, age int, email string) (int64, error) {

	//    s.Db.Prepare("I")

	stmt, err := s.Db.Prepare("Insert into Student(Name,Age,Email) values(?,?,?)")
	if err != nil {

		fmt.Println("Error creating Statement for inserting Student")
		return 0, err
	}
	defer stmt.Close()
	Result, err := stmt.Exec(name, age, email)

	if err != nil {
		fmt.Println("Error executing insert query for Student")
		return 0, err
	}

	lastid, _ := Result.LastInsertId()

	return lastid, nil

}

func (s *Sqlite) GetStudent(id int64) (Type.Student, error) {

	stmt, err := s.Db.Prepare("Select * from Student where ID =? ")

	if err != nil {
		fmt.Println("Error creating a statement for getting student")
	}

	defer stmt.Close()
	var studentdata = Type.Student{}
	err = stmt.QueryRow(id).Scan(&studentdata.Id, &studentdata.Name, &studentdata.Age, &studentdata.Email)
	// Result,err:=stmt.Exec(id)

	if err != nil {

		return studentdata, fmt.Errorf("error executing the select querry for the given id,%d", id)

	}
	return studentdata, nil

}

func (s *Sqlite) ListAllStudents() ([]Type.Student, error) {

	stmt, err := s.Db.Prepare("Select * from Student")

	if err != nil {
		return []Type.Student{}, fmt.Errorf("Error Preparing statement for Listing all Students")
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {

		return []Type.Student{}, fmt.Errorf("Error Executing Query for Listing all Students")
	}

	var students = []Type.Student{}
	defer rows.Close()
	for rows.Next() {

		var student Type.Student

		rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email)
		students = append(students, student)
	}
	return students, nil
}

func (s *Sqlite) UpdateStudent(id int64, name string, age int, email string) (int64, error) {

	stmt, err := s.Db.Prepare("Update Student set Name =? ,Age=?,Email=? where Id=?")
	if err != nil {
		return 0, fmt.Errorf("error preparing statement for updating student with id %d", id)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, age, email, id)

	if err != nil {
		return 0, fmt.Errorf("error executing update query for student with id %d", id)
	}

	return id, nil

}

func (s *Sqlite) DeleteStudent(id int64) (int64, error) {

	stmt, err := s.Db.Prepare("Delete from Student where Id =?")

	if err != nil {
		return 0, fmt.Errorf("error preparing the statement to delete the student id %d", id)
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		return 0, fmt.Errorf("error Executing the Query to delete the student id %d", id)
	}

	return id, nil

}
