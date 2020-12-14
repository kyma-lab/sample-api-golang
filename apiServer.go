package main

import (
	"fmt"
	BookApi "maihiro.com/sampleApi/server/booksApi"
)

func main() {
	var mybooks []BookApi.Book

	mybooks = append(mybooks, BookApi.Book{ID: "1", Isbn: "213333", Title: "Mein Buch", Author: &BookApi.Author{Firstname: "Dr.", Lasname: "no"}})
	mybooks = append(mybooks, BookApi.Book{ID: "2", Isbn: "213333", Title: "Das Telefonbuch", Author: &BookApi.Author{Firstname: "Die.", Lasname: "Telekom"}})

	fmt.Printf("Starting Sample Api Server with  %d dummy elements", len(mybooks))

	BookApi.InitBookApi(mybooks)

}
