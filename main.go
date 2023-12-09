package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("API started ...")

	handlerLogin := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	http.HandleFunc("/", handlerLogin)
	http.HandleFunc("/login", handlerLogin)

	http.ListenAndServe(":3001", nil)
}
