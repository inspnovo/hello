package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello!")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
	fmt.Println("hello,vim-go")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
