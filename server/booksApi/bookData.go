package books

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jucardi/go-streams/streams"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lasname   string `json:"lastname"`
}

func InitBookApi(initialStock []Book) {
	fmt.Println(".. intitialize server")
	books = initialStock
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	element := streams.
		FromArray(books).
		Filter(func(book interface{}) bool {
			return strings.EqualFold(book.(Book).ID, params["id"])
		}).
		First("empty")

	if element != nil {
		json.NewEncoder(w).Encode(element)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = "mb-" + strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = removeBookPerIndex(books, index)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBookPerIndex(books []Book, index int) []Book {
	return append(books[:index], books[index+1:]...)
}
