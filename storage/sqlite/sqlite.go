package sqlite

import (
	"database/sql"
	"fmt"
	"myapp/config"

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
	Result, err := stmt.Exec(name, age, email)

	if err != nil {
		fmt.Println("Error executing insert query for Student")
		return 0, err
	}

	lastid, _ := Result.LastInsertId()

	return lastid, nil

}
