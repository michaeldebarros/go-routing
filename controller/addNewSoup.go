package controller

import (
	"router/model"
	"strings"
)

//AddNewSoup exported
func AddNewSoup(name []string, origin []string, ingredients []string, spicy []string, success chan bool) {
	session := MgoSession.Copy()
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

	if len(spicy) == 0 {
		spiceFactor = false
	} else {
		spiceFactor = true
	}

	//create instance in Soup struct
	s := model.Soup{
		Name:        name[0],
		Origin:      origin[0],
		Spicy:       spiceFactor,
		Ingredients: ingrSlice2,
	}
	//insert the instance
	if err := c.Insert(s); err != nil {
		panic(err)
	}
	success <- true
}
