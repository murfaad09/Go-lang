package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var logins = []login{
	{Username: "Bentley", Password: "Red"},
	{Username: "Mercedez", Password: "Black"},
	{Username: "Jaguar", Password: "Silver"},
}

func loginCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username1 := vars["Username"]
	password1 := vars["Password"]
	for _, c := range logins {
		if c.Username == username1 && c.Password == password1 {
			//fmt.Println("Suucessfuly Login", c)
			return
		}

	}

}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/logins/", loginCheck).Methods("POST")
	http.ListenAndServe(":8080", r)
	//r.HandleFunc("/logins/", func(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Login Successfull",)
}
