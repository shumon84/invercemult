package main

import (
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/post.html")
}
