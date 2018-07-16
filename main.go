package main

import (
	"flag"
	"net/http"
)

func main() {
	// オプションをパース
	var (
		port string
		help bool
	)
	flag.StringVar(&port, "p", "8080", "割り当てるポート")
	flag.BoolVar(&help, "h", false, "ヘルプを表示")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	server := http.Server{
		Addr: ":" + port,
	}
	http.HandleFunc("/color", color)
	http.HandleFunc("/post", post)
	http.HandleFunc("/", post)
	server.ListenAndServe()
}
