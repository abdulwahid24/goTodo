package main

import (
	"os"
	"log"
	"net/http"
)


func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), corsHandler(router))) 

}
