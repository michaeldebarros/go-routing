package usersession

import "net/http"

//MakeCookie export function
func MakeCookie(userIDHex string) *http.Cookie {
	newCookie := http.Cookie{
		Name:   "soup-site-userID",
		Value:  userIDHex,
		MaxAge: 10,
	}
	return &newCookie
}

//DELETE COOKIE
