package controller

import (
	"fmt"
	"router/model"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2/bson"
)

//UserLogin export
func UserLogin(login []string, password []string, message chan string) {
	session := MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("users")

	userByLogin := model.User{}

	c.Find(bson.M{"login": login[0]}).One(&userByLogin)

	//if there is no user in the db with that login
	if len(userByLogin.ID) == 0 {
		//create new user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password[0]), 4)
		if err != nil {
			fmt.Println(err)
		}
		//insert new user in db
		userToBeInserted := model.User{
			Login:    login[0],
			Password: hashedPassword,
		}
		if err := c.Insert(userToBeInserted); err != nil {
			fmt.Println(err)
		}
		message <- "New user created"
	} else {
		//if there already is a user in the db make login
		err := bcrypt.CompareHashAndPassword(userByLogin.Password, []byte(password[0]))
		if err != nil {
			fmt.Println(err)
		}
		message <- "User logged in successfully"
	}
}
