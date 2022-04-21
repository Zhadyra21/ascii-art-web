package main

import (
	"log"
	"net/http"
	asciiartweb "web/web-ascii-art"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	asciiartweb.Init()
	log.Println("\nThe server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
