package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Student struct {
	Name   string
	Roll   string
	Gender string
}

func main() {

	//student := Student{"anshul", "2", "male"}
	var db *sql.DB
	db, err := sql.Open("postgres", "user=postgres password=redhat dbname=testdb host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		err = errors.Wrapf(err, "database not connected ")
		return
	}
	fmt.Println("database connected;")
	/*
		_, tableCreateError := db.Exec(`CREATE TABLE student(
			NAME	TEXT	NOT NULL,
			ROLL	TEXT	NOT NULL,
			GENDER	TEXT	NOT NULL
		)`)
		if tableCreateError != nil {
			errors.Wrapf(tableCreateError, "Table crate err")

		}
		fmt.Println("Error:", tableCreateError)
	*/
	//fmt.Println("type of table", reflect.TypeOf(table), *table)
	students := GetStudents(db)
	fmt.Println("Student:- ", students[0].Name)
	defer db.Close()
	fmt.Println("database closed successfully")
}

func GetStudents(db *sql.DB) []Student {

	rows, testerr := db.Query("SELECT * FROM student")
	if testerr != nil {
		fmt.Println("Select Query Error", testerr)
	}
	var students []Student
	for rows.Next() {
		var student = Student{}
		err := rows.Scan(&student.Name, &student.Roll, &student.Gender)
		if err != nil {
			fmt.Println("Read Error", err)
		}
		students = append(students, student)
	}
	return students
}

/*
func InsertStudent(db *sql.DB, student Student) (err error) {

}
*/
