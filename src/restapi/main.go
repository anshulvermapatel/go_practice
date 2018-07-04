package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//"fmt"
	//"http"
	//"encoding/json"
)

type Student struct {
	Rollno    int64   `json:rollno`
	Name      string  `json:name`
	Sex       string  `json:sex`
	Aggregate float32 `json:aggregate`
}

func (s Student) ToJSON() []byte {
	JSON, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return JSON
}

func main() {

	//name := "anshul"
	http.HandleFunc("/api/details/student", WelcomeFunc)
	http.HandleFunc("/api/details/student/", DetailByRollNo)
	//http.HandleFunc("/api/iterate", IterateContent)
	http.ListenAndServe(":8000", nil)
}

var students []Student = []Student{{10, "anshul", "male", 9.5}, {11, "jassi", "male", 8.9}, {12, "Ayushi", "female", 8.5}}

func WelcomeFunc(w http.ResponseWriter, r *http.Request) {
	//Students, _ := json.Marshal(s)
	//writeJSON(w, s)
	//w.Header().Set("Content-Type", "application/json")
	l := r.URL.Path[len("/api/details/student/"):]
	fmt.Fprintf(w, l)
}

func DetailByRollNo(w http.ResponseWriter, r *http.Request) {
	studentRollNo := r.URL.Path[len("/api/details/student/"):]

	//fmt.Fprintf(w, string(rollno))
	for _, student := range students {
		rollno, _ := strconv.ParseInt(studentRollNo, 10, 32)
		if student.Rollno == rollno {
			studentJSON, _ := json.Marshal(student)
			w.Header().Add("Content-Type", "application/json")
			w.Write(studentJSON)
			return
		}
	}
	//w.Header().Add("Content-Type", "text/plain")
	//w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, studentRollNo+" not found")
}

/*
func IterateContent(w http.ResponseWriter, r *http.Request) {
	content := regexp.MustCompile("/").Split(r.URL.Path, 2)
	//for _, con := range content {
	fmt.Fprintf(w, content[1])
	//}
}
*/
