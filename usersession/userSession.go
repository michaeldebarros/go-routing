package usersession

import (
	"fmt"
	"net/http"
	"router/db"
	"router/model"

	"gopkg.in/mgo.v2/bson"
)

//create the session MAP

//SessionMAP export
var SessionMAP map[string]string

func init() {
	m := make(map[string]string)
	SessionMAP = m
}

//InitSession export function
func InitSession(userIDHex string) *http.Cookie {
	//creat a session and insert in db
	session := db.MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("sessions")

	s := model.Session{
		ID:     bson.NewObjectId(),
		UserID: userIDHex,
	}

	if err := c.Insert(s); err != nil {
		fmt.Println(err)
	}

	sessionIDHex := s.ID.Hex()

	//write session to SessionMAP
	SessionMAP[sessionIDHex] = userIDHex

	fmt.Println(SessionMAP)

	//create cookie and return it
	newCookie := http.Cookie{
		Name:   "session",
		Value:  sessionIDHex,
		MaxAge: 60,
	}

	return &newCookie
}

//DeleteSession export
func DeleteSession(sessionIDString string, success chan bool) {
	//creat a session and insert in db
	session := db.MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("sessions")

	if err := c.RemoveId(bson.ObjectIdHex(sessionIDString)); err != nil {
		fmt.Println(err)
	}

	//delete from SessionMAP
	delete(SessionMAP, sessionIDString)

	fmt.Println(SessionMAP)

	success <- true

}