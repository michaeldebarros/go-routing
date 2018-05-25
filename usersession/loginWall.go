package usersession

import (
	"net/http"
	"strings"
)

//LoginWall export
func LoginWall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check all routes besides login and assets if cookie is present

		if r.URL.String() != "/login" && !strings.HasPrefix(r.URL.String(), "/assets") {
			cookie, err := r.Cookie("session")
			if err != nil {
				http.Redirect(w, r, "/login", 302)
				return
			}
			_, ok := SessionMAP[cookie.Value]
			if !ok {
				http.Redirect(w, r, "/login", 302)
			}
		}
		next.ServeHTTP(w, r)
	})
}

//work on this here to make the loginwall
