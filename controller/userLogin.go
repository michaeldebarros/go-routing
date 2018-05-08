package controller

import (
	"fmt"
	"router/model"

	"gopkg.in/mgo.v2/bson"
)

//UserLogin export
func UserLogin(login []string, password []string) {
	session := MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("users")

	queryResultByLogin := model.User{}
	c.Find(bson.M{"login": login[0]}).One(&queryResultByLogin)
	fmt.Println(queryResultByLogin)
}
