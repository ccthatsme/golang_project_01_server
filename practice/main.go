package main

import (
 "fmt"
// "encoding/json"
"log"
"net/http"
// "math/rand"
// "strconv"
"github.com/gorilla/mux"

)



func main() {

//init router
r:=mux.NewRouter()

//route handlers/endpoints
r.HandleFunc("/api/books", getBooks).Methods("GET")
// r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
// r.HandleFunc("/api/books", createBook).Methods("POST")
// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

//similar to express listen
//r is the handler to handle requests get post etc.
log.Fatal(http.ListenAndServe(":8000", r))
}


func getBooks(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "<h1>Hellow World<h1>")
}


