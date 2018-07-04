package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book {
	var book Book
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

func main() {
	http.HandleFunc("/getjson", getjson)
	http.ListenAndServe(":8080", nil)

	book := Book{Title: "Book1", Author: "anshul", ISBN: "12223"}
	jsondata := book.ToJSON()
	//fmt.Println(string(jsondata))
	book = FromJSON(jsondata)
	//fmt.Println(book)
}

func getjson(w http.ResponseWriter, r *http.Request) {
	book := Book[{Title: "Book1", Author: "anshul", ISBN: "12223"},{Title: "Book2", Author: "khushal", ISBN: "123674567"}]
	bytes := book.ToJSON()
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(bytes)
}
