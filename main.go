package main

import (
	"io"
	"net/http"
)

func content() string {
	return "Hello world!"
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, content())
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
