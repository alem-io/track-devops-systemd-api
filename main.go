package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = "9090"
)

var (
	inc = 0
)

func root(w http.ResponseWriter, req *http.Request) {
	resp := fmt.Sprintf("web-server: %d\n", inc)
	fmt.Fprint(w, resp)
	inc++
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":"+PORT, nil)
}
