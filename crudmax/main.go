package main

import "github.com/gorilla/mux"

func main() 
{
	 router := mux.NeNewRouter().StrictSlash(true)
	 router.Handlefunc("/cars", returnAllCArs.Methods("Get") )
}