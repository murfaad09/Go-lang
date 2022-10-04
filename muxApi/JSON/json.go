// package main
// import (
// 	"fmt"
// 	"net/http"
// 	"encoding/json"
// )
// type User struct{
// 	FName string `json:"fname"`
// 	LName string `json:"lname"`
// 	Age int `json:"age"`
// }
// 		func main(){
// 			http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request){
// 				var user User
// 				json.NewDecoder(r.Body).Decode(&user)
// 				fmt.Printf("&s &s is &s years old", user.Fname,user.Lname,user.Age)
// 				})

// 				http.HandleFunc("/encode",func (w http.ResponseWriter, r *http.){
// 					piter:= &User{
// 						Fname: "JHon"
// 						Lname: "Lee"
// 						Age: 32
// 						json.NewEncoder(w).Encode(piter)

// 			}
// 			})
				
// 					http.ListenAndServe(":8081",nil)
				
// 				}	

// json.go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}

func main() {
    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)

        fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
    })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        peter := User{
            Firstname: "John",
            Lastname:  "Doe",
            Age:       25,
        }

        json.NewEncoder(w).Encode(peter)
    })

    http.ListenAndServe(":8080", nil)
}