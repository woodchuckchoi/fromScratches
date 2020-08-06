package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println(w.Header())
	fmt.Fprintf(w, "hi\n")
}

func main() {
	http.HandleFunc("/", test)

	http.ListenAndServe(":7777", nil)
}
