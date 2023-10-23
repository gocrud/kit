package mgo

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringId() string {
	return primitive.NewObjectID().Hex()
}

func ObjectId() primitive.ObjectID {
	return primitive.NewObjectID()
}

type ObjectID = primitive.ObjectID
