package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Text struct {
	WindowsName string             `json:"windows_name" bson:"windows_name"`
	Keystrokes  string             `json:"keystrokes" bson:"keystrokes"`
	Timestamp   primitive.DateTime `json:"timestamp" bson:"timestamp"`
}
