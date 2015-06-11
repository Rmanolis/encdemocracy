package models

import "gopkg.in/mgo.v2/bson"

type Vote struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	IsAccepted  bool
	PublicKeyId bson.ObjectId
	TopicId     bson.ObjectId
}
