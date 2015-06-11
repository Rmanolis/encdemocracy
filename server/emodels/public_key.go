package models

import "gopkg.in/mgo.v2/bson"

type PublicKey struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	JSON     string
	CheckSum string
}
