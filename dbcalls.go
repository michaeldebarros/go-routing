package main

import (
	"strings"

	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	//	defer session.Close()  // take a closer look at this
	mgoSession = session
}

func addNewSoup(name []string, origin []string, ingredients []string, spicy []string) {
	session := mgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("soups")

	//parse ingredients
	ingrSlice1 := strings.Split(ingredients[0], ",")
	var ingrSlice2 []string
	for _, individualIngredient := range ingrSlice1 {
		trimmedIngretdient := strings.TrimSpace(individualIngredient)
		ingrSlice2 = append(ingrSlice2, trimmedIngretdient)
	}

	//parse spicy
	var spiceFactor bool
	if spicy[0] == "true" {
		spiceFactor = true
	} else {
		spiceFactor = false
	}

	//create instance in Soup struct
	s := Soup{
		Name:        name[0],
		Origin:      origin[0],
		Spicy:       spiceFactor,
		Ingredients: ingrSlice2,
	}
	//insert the instance
	if err := c.Insert(s); err != nil {
		panic(err)
	}
}
