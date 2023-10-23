package mgo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Filter struct {
	tbl    *mongo.Collection
	filter bson.M
	ctx    context.Context
	sort   any
	skip   int64
	limit  int64
}

func (f *Filter) One(data any) error {
	opts := options.FindOne()
	if f.sort != nil {
		opts.SetSort(f.sort)
	}
	if f.skip > 0 {
		opts.SetSkip(f.skip)
	}
	err := f.tbl.FindOne(f.ctx, f.filter, opts).Decode(data)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	return err
}

func (f *Filter) List(data any) error {
	opts := options.Find()
	if f.sort != nil {
		opts.SetSort(f.sort)
	}
	if f.skip > 0 {
		opts.SetSkip(f.skip)
	}
	if f.limit > 0 {
		opts.SetLimit(f.limit)
	}
	cur, err := f.tbl.Find(f.ctx, f.filter, opts)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	if err != nil {
		return err
	}
	defer cur.Close(f.ctx)
	return cur.All(f.ctx, data)
}

func (f *Filter) Count() (int64, error) {
	return f.tbl.CountDocuments(f.ctx, f.filter)
}

func (f *Filter) Sort(sort any) *Filter {
	f.sort = sort
	return f
}

func (f *Filter) Skip(skip int64) *Filter {
	f.skip = skip
	return f
}

func (f *Filter) Limit(limit int64) *Filter {
	f.limit = limit
	return f
}

func (f *Filter) Scopes(scopes ...ScopeFilter) *Filter {
	for _, scope := range scopes {
		scope(f)
	}
	return f
}

func (f *Filter) Filter(key string, value any) *Filter {
	f.filter[key] = value
	return f
}

func (f *Filter) Filters(filters bson.M) *Filter {
	for key, value := range filters {
		f.filter[key] = value
	}
	return f
}

func (f *Filter) UpdateOne(update any) error {
	_, err := f.tbl.UpdateOne(f.ctx, f.filter, update)
	return err
}

func (f *Filter) UpdateMany(update any) error {
	_, err := f.tbl.UpdateMany(f.ctx, f.filter, update)
	return err
}

func (f *Filter) DeleteOne() error {
	_, err := f.tbl.DeleteOne(f.ctx, f.filter)
	return err
}

func (f *Filter) DeleteMany() error {
	_, err := f.tbl.DeleteMany(f.ctx, f.filter)
	return err
}
