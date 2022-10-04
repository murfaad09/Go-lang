package main

import (
	"net/http"
)

// func helloworld(w http.ResponseWriter, r *http.Request) {

// }

// func main() {
// 	http.HandleFunc("/murfaad", helloworld)
// 	if err := http.ListenAndServe("localhost:8081", nil); err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// 	// if err!= nil{
// 	// 	panic(err)
// 	// }
// }

// func main() {
// 	http.HandleFunc("/murfaad", helloworld)
// 	err := http.ListenAndServe(":8081", nil)
// 	if err != nil {
// 		panic(err)
// }
// }
func main() {
	http.HandleFunc("/murfaad", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello WOrld"))
	})
	http.ListenAndServe(":8081", nil)
}
