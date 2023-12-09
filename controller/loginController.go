package controller

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login page requested ...")
	w.Write([]byte("this is login page"))
}
