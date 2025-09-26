//go:build release
package main

import (
	"embed"
	"net/http"
)

//go:embed assets/* web/*
var embeddedFiles embed.FS

func SetupFileServer(mux *http.ServeMux) {
	mux.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.FS(embeddedFiles)),
		),
	)

	mux.Handle("/web/",
		http.StripPrefix("/web/",
			http.FileServer(http.FS(embeddedFiles)),
		),
	)
}
