package main

import (
	"log"
	"net/http"
	"time"
)

const PORT = ":6969"

func main() {
	mux := http.NewServeMux()
  SetupFileServer(mux)

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
