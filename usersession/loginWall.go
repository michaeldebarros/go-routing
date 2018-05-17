package usersession

import (
	"net/http"
)

//LoginWall export
func LoginWall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check all routes besides login and static if cookie is present
		if r.URL.String() != "/login" && r.URL.String() != "/static" {
			cookie, err := r.Cookie("session")
			if err != nil {
				http.Redirect(w, r, "/login", 302) //fix this redirect
				return
			}
			_, ok := SessionMAP[cookie.Value]
			if !ok {
				http.Redirect(w, r, "/login", 302) //fix this redirect
			}
		}
		next.ServeHTTP(w, r)
	})
}

//work on this here to make the loginwall
