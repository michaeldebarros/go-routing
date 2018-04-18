package main

import mgo "gopkg.in/mgo.v2"

var mgoSession *mgo.Session

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	mgoSession = session
}

func addNewSoup(name []string, origin []string, ingredients []string, spicy []string) {
	session := mgoSession.Copy()
	defer session.Close()
}
