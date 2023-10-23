package mgo

type ScopeFilter func(f *Filter)
type ScopeAggregate func(agg *Aggregate) *Aggregate
