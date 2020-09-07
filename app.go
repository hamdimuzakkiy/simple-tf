package main

import (
	"net/http"

	"github.com/hamdimuzakkiy/simple-tf/src/handler"
)

func main() {
	http.HandleFunc("/", handler.MainHandler)

	http.ListenAndServe(":9119", nil)
}
