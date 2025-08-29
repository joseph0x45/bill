package main

import (
	"log"
	"net/http"
	"time"
)

const PORT = ":6969"

func main() {
	mux := http.NewServeMux()

  mux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	log.Println("[INFO]: Server listening on port ", PORT)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
