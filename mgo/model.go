package mgo

type Model interface {
	TableName() string
}
