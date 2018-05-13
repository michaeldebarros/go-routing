package usersession

import "net/http"

//MakeCookie export
func MakeCookie(userIDHex string) *http.Cookie {
	newCookie := http.Cookie{
		Name:   "soup-site-userID",
		Value:  userIDHex,
		MaxAge: 10,
	}
	return &newCookie
}
