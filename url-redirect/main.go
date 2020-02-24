package main

import (
	"fmt"
	"github.com/sathishkumar-manogaran/GoLangPrograms/url-redirect/url-handler"
	"net/http"

	//"github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/facebook-page": "https://facebook.com",
		"/twitter-page":     "https://twitter.com",
	}
	mapHandler := url_handler.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml :=
		`
- path: /github
  url: https://github.com
- path: /gitlab
  url: https://gitlab.com
`
	yamlHandler, err := url_handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
