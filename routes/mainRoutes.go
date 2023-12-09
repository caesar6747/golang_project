package routes

import "net/http"

func Register() {
	http.HandleFunc("/login", controller.loginHandler)
}
