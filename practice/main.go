package main

import (
 "fmt"
 "encoding/json"
"log"
"net/http"
// "math/rand"
 "strconv"
"github.com/gorilla/mux"

)

//Book Struct (Model)
type Book struct{

//json field that we fetch
ID string `json:"id"`
Isbn string `json:"isbn"`
Title string `json:"title"`
Author *Author `json:"author"`

}

//Author struct
type Author struct {
Firstame string `json:firstname`
Lastname string `json:lastname`
}

//init books var as a slice book struct
var books []Book



func main() {

//init router
r:=mux.NewRouter()

//mock data
books = append(books, Book{ID: "1", Isbn:"448743", Title: "Book ONe", Author: &Author{Firstame:"Chris", Lastname:"Ciric"}})
books = append(books, Book{ID: "2", Isbn:"745367", Title: "Book Two", Author: &Author{Firstame:"josh", Lastname:"johnson"}})


//route handlers/endpoints
r.HandleFunc("/api/books", getBooks).Methods("GET")
r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
 r.HandleFunc("/api/books", createBook).Methods("POST")
// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

//similar to express listen
//r is the handler to handle requests get post etc.
log.Fatal(http.ListenAndServe(":8000", r))
}


//get all books
func getBooks(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(books)
//  fmt.Fprintf(w, "<h1>Hellow World<h1>")
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
params := mux.Vars(r) //Get Params

for _,item := range books {
if item.ID == params["id"]{
json.NewEncoder(w).Encode(item)
return
}
}
json.NewEncoder(w).Encode(&Book{})
}

//post a book
func createBook(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")

var book Book
_ = json.NewDecoder(r.Body).Decode(&book)
book.ID = strconv.Itoa(3)

books = append(books, book)
json.NewEncoder(w).Encode(book)
}

//update a book
func updateBook(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "<h1>Hellow World<h1>")
}

//delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "<h1>Hellow World<h1>")
}

