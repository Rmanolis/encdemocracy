package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Email string
	Pass  string
}
