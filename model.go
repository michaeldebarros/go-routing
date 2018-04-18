package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Soup export
type Soup struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	Origin      string        `bson:"origin" json:"origin"`
	Spicy       bool          `bson:"spicy" json:"spicy"`
	Ingredients []string      `bson:"ingredients" json:"ingredients"`
	TimeStamp   time.Time     `bson:"timeStamp" json:"timeStamp"`
}
