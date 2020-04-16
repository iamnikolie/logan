package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type LogMessage struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Service   string        `json:"service" bson:"service"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	Level     string        `json:"level" bson:"level"`
	Message   string        `json:"message" bson:"message"`
	Tags      []string      `json:"tags" bson:"tags"`
}
