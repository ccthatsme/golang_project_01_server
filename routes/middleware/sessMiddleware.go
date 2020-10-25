package middleware

import (
	"fmt"
	"net/http"

)

func checkTokenExists(next http.HandlerFunc) http.HandlerFunc{
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
ctx := r.Context()
fmt.Println(ctx)
})
}