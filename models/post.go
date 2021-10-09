package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Posts struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Caption         string        `json:"caption" bson:"caption"`
	ImageUrl        string        `json:"imageurl" bson:"imageurl"`
	UserId          string        `json:"userid" bson:"userid"`
	PostedTimeStamp time.Time     `json:"postedTimestamp"`
}
