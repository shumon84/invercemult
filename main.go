package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":9000",
	}
	http.HandleFunc("/color", color)
	http.HandleFunc("/post", post)
	http.HandleFunc("/", post)
	server.ListenAndServe()
}
