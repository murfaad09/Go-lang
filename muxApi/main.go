package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w  http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	http.ListenAndServe(":8080", r)
}
func ReadBook(w  http.ResponseWriter, r *http.Request){
	fmt.Println("Hi RAEd Book")
}