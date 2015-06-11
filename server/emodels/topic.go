package models

import "gopkg.in/mgo.v2/bson"

type Topic struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Title  string
	Text   string
	UserId bson.ObjectId
}
