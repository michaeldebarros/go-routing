package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func deleteSoup(id []string, success chan bool) {
	session := mgoSession.Copy()
	defer session.Close()

	fmt.Println(bson.IsObjectIdHex(id[0]))
	c := session.DB("RECEPIES").C("soups")
	err := c.RemoveId(bson.ObjectIdHex(id[0]))
	if err != nil {
		fmt.Println(err)
	}
	success <- true
}
