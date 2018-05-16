package usersession

import (
	"fmt"
	"net/http"
)

//LoginWall export
func LoginWall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cookie)
		next.ServeHTTP(w, r)
	})
}
