package controller

import (
	"fmt"
	"router/model"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2/bson"
)

//UserLogin export
func UserLogin(login []string, password []string) {
	session := MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("users")

	userByLogin := model.User{}

	c.Find(bson.M{"login": login[0]}).One(&userByLogin)
	fmt.Println(len(userByLogin.ID))

	if len(userByLogin.ID) == 0 {
		//create new user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password[0]), 4)
		if err != nil {
			fmt.Println("Error while hashing password")
		}
		//insert new user in db
		userToBeInserted := model.User{
			Login:    login[0],
			Password: hashedPassword,
		}
		if err := c.Insert(userToBeInserted); err != nil {
			fmt.Println("problem inserting user")
		}
	}
}
