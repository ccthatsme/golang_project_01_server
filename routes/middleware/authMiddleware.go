package middleware

import (
	"fmt"
	"net/http"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
		fmt.Println("ok: ", ok)
		if !ok || !checkUserAndPassword(username, password) {
			w.Header().Set("x-auth-token", "invalid")
			w.WriteHeader(401)
			w.Write([]byte("not authorized"))
			handler.ServeHTTP(w, r)
			return
		} else {
			handler.ServeHTTP(w, r)
		}

	})
}

func checkUserAndPassword(username, password string) bool {
	return username == "cc" && password == "password"
}
