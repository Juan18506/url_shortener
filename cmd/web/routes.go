package main

import (
	"fmt"
	"net/http"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
