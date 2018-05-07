package controller

import "router/model"

//GetAllSoups export
func GetAllSoups() ([]model.Soup, error) {
	session := MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("soups")

	//query the db for all soups
	results := []model.Soup{}
	err := c.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}
	return results, err
}
