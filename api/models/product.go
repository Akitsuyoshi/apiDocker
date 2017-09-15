package models

import (
	"labix.org/v2/mgo/bson"
)

// Product explain the data model struct in this application.
type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `form:"title" json:"title"`
	Description string        `form:"description" json:"description"`
	Price       float64       `form:"price" json:"price"`
}
