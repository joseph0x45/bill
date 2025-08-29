package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

  mux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Addr:         ":6969",
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	log.Println("[INFO]: Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
