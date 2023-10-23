package mgo

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	cli *Client
)

func init() {
	var err error
	cli, err = NewClient(context.TODO(), Options{
		Uri: "mongodb://root:root@127.0.0.1/?directConnection=true",
	})
	if err != nil {
		panic(err)
	}

}

type User struct {
	Id   ObjectID `bson:"_id"`
	Name string
	Age  int
}

func TestInsertOne(t *testing.T) {
	db := cli.Database("test")
	tbl := db.Table("user")
	i := ObjectId()
	t.Log(i)
	var u = User{
		Id:   i,
		Name: "test",
		Age:  18,
	}
	id, err := tbl.InsertOne(context.Background(), u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestInsertMany(t *testing.T) {
	db := cli.Database("test")
	tbl := db.Table("user")
	var u = []any{
		User{
			Name: "test1",
			Age:  16,
		},
		User{

			Name: "test2",
			Age:  20,
		},
		User{
			Name: "test3",
			Age:  22,
		},
		User{
			Name: "test4",
			Age:  35,
		},
	}
	_, err := tbl.InsertMany(context.Background(), u)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilter(t *testing.T) {
	db := cli.Database("test")
	tbl := db.Table("user")
	var u = User{}
	err := tbl.Filter(context.Background()).One(&u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestFilterAll(t *testing.T) {
	db := cli.Database("test")
	tbl := db.Table("user")
	var u = []User{}
	err := tbl.Filter(context.Background()).
		Sort(bson.M{"age": -1}).
		Skip(1).
		Limit(2).
		List(&u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestFilterCount(t *testing.T) {
	db := cli.Database("test")
	tbl := db.Table("user")
	count, err := tbl.Filter(context.Background()).Filter("age", bson.M{"$gt": 20}).Count()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("count", count)
}

func TestAggregate_List(t *testing.T) {
	db := cli.Database("test")
	tbl := db.Table("user")
	var u = []User{}
	err := tbl.Aggregate(context.Background()).Match(bson.M{"age": bson.M{"$lt": 20}}).List(&u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
