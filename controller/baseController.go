package controller

import "net/http"

func testController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is test page"))
}
