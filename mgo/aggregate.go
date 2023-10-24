package mgo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Aggregate struct {
	tbl   *mongo.Collection
	ctx   context.Context
	pipes bson.A
}

func (a *Aggregate) List(data any) error {
	cur, err := a.tbl.Aggregate(a.ctx, a.pipes)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	if err != nil {
		return err
	}
	defer cur.Close(a.ctx)
	err = cur.All(a.ctx, data)
	return err
}

func (a *Aggregate) One(data any) error {
	cur, err := a.tbl.Aggregate(a.ctx, a.pipes)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	if err != nil {
		return err
	}
	defer cur.Close(a.ctx)
	if cur.TryNext(a.ctx) {
		err = cur.Decode(data)
	}
	return err
}

func (a *Aggregate) Scopes(scopes ...ScopeAggregate) *Aggregate {
	for _, scope := range scopes {
		scope(a)
	}
	return a
}

func (a *Aggregate) AddFields(addFields bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$addFields", Value: addFields}})
	return a
}

func (a *Aggregate) Bucket(bucket bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$bucket", Value: bucket}})
	return a
}

func (a *Aggregate) BucketAuto(bucketAuto bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$bucketAuto", Value: bucketAuto}})
	return a
}

func (a *Aggregate) CollStats(collStats bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$collStats", Value: collStats}})
	return a
}

func (a *Aggregate) Count(field string) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$count", Value: field}})
	return a
}

func (a *Aggregate) Densify(densify bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$densify", Value: densify}})
	return a
}

func (a *Aggregate) Facet(facet bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$facet", Value: facet}})
	return a
}

func (a *Aggregate) Fill(fill bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$fill", Value: fill}})
	return a
}

func (a *Aggregate) GeoNear(geoNear bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$geoNear", Value: geoNear}})
	return a
}

func (a *Aggregate) GraphLookup(graphLookup bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$graphLookup", Value: graphLookup}})
	return a
}

func (a *Aggregate) Group(group bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$group", Value: group}})
	return a
}

func (a *Aggregate) IndexStats(indexStats bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$indexStats", Value: indexStats}})
	return a
}

func (a *Aggregate) Limit(limit int64) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$limit", Value: limit}})
	return a
}

func (a *Aggregate) Lookup(lookup bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$lookup", Value: lookup}})
	return a
}

func (a *Aggregate) Match(match bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$match", Value: match}})
	return a
}

func (a *Aggregate) Merge(merge bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$merge", Value: merge}})
	return a
}

func (a *Aggregate) Out(out bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$out", Value: out}})
	return a
}

func (a *Aggregate) Project(project bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$project", Value: project}})
	return a
}

func (a *Aggregate) Redact(redact bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$redact", Value: redact}})
	return a
}

func (a *Aggregate) ReplaceRoot(replaceRoot bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$replaceRoot", Value: replaceRoot}})
	return a
}

func (a *Aggregate) ReplaceWith(replaceWith bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$replaceWith", Value: replaceWith}})
	return a
}

func (a *Aggregate) Sample(sample bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$sample", Value: sample}})
	return a
}

func (a *Aggregate) Search(search bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$search", Value: search}})
	return a
}

func (a *Aggregate) SearchMeta(searchMeta bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$searchMeta", Value: searchMeta}})
	return a
}

func (a *Aggregate) Set(set bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$set", Value: set}})
	return a
}

func (a *Aggregate) SetWindowFields(setWindowFields bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$setWindowFields", Value: setWindowFields}})
	return a
}

func (a *Aggregate) Skip(skip int64) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$skip", Value: skip}})
	return a
}

func (a *Aggregate) Sort(sort bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$sort", Value: sort}})
	return a
}

func (a *Aggregate) SortByCount(sortByCount bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$sortByCount", Value: sortByCount}})
	return a
}

func (a *Aggregate) UnionWith(unionWith bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$unionWith", Value: unionWith}})
	return a
}

func (a *Aggregate) Unset(unset bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$unset", Value: unset}})
	return a
}

func (a *Aggregate) Unwind(unwind bson.M) *Aggregate {
	a.pipes = append(a.pipes, bson.D{{Key: "$unwind", Value: unwind}})
	return a
}
