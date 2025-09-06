//go:build !release
package main

import (
	"net/http"
)

func SetupFileServer(mux *http.ServeMux) {
	mux.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("./assets")),
		),
	)

	mux.Handle("/web/",
		http.StripPrefix("/web/",
			http.FileServer(http.Dir("./web")),
		),
	)
}
