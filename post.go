package main

import (
	"fmt"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w, r, "html/post.html")
}
