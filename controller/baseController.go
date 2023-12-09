package controller

import "net/http"

func TestController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is test page"))
}
