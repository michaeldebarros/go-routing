package main

func getAllSoups() ([]Soup, error) {
	session := mgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("soups")

	//query the db for all soups
	results := []Soup{}
	err := c.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}
	return results, err
}
