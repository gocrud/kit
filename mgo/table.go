package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Table struct {
	tbl *mongo.Collection
}

func NewTable(db *Database, name string) *Table {
	return &Table{tbl: db.db.Collection(name)}
}

func (t *Table) Filter(ctx context.Context) *Filter {
	return &Filter{tbl: t.tbl, ctx: ctx, filter: make(bson.M)}
}

func (t *Table) InsertOne(ctx context.Context, data any) (id any, err error) {
	var result *mongo.InsertOneResult
	result, err = t.tbl.InsertOne(ctx, data)
	if err != nil {
		return
	}
	id = result.InsertedID
	return
}

func (t *Table) InsertMany(ctx context.Context, data []any) (ids []any, err error) {
	var result *mongo.InsertManyResult
	result, err = t.tbl.InsertMany(ctx, data)
	if err != nil {
		return
	}
	ids = result.InsertedIDs
	return
}

func (t *Table) Drop(ctx context.Context) (err error) {
	err = t.tbl.Drop(ctx)
	return
}

func (t *Table) Aggregate(ctx context.Context) *Aggregate {
	return &Aggregate{tbl: t.tbl, ctx: ctx}
}
