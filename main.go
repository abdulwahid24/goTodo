package main

import (
	"log"
	"net/http"
)


func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe("https://secure-brook-4935.herokuapp.com", corsHandler(router))) 

}
