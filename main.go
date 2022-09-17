package main

import (
	"log"
	"net/http"
)

func main() {
	// create http file system
	fs := http.FileServer(http.Dir("./static"))

	// start listening and root
	http.Handle("/", fs)

	// starting server
	log.Println("Starting server on :3000 port ...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf(">>> error serving %v", err.Error())
	}

}