package mgo

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	db *mongo.Database
}

func NewDatabase(cli *Client, name string) *Database {
	return &Database{db: cli.cli.Database(name)}
}

func (d *Database) Table(name string) *Table {
	return &Table{tbl: d.db.Collection(name)}
}

func (d *Database) Model(m Model) *Table {
	return d.Table(m.TableName())
}
