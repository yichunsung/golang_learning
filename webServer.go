package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hllo World! Golang Web Server"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8081", nil)
}

