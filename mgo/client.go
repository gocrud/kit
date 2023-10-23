package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	cli *mongo.Client
}

type Options struct {
	Uri string
}

func NewClient(ctx context.Context, opts Options) (*Client, error) {
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(opts.Uri))
	if err != nil {
		return nil, err
	}
	return &Client{cli: cli}, nil
}

func (c *Client) Database(name string) *Database {
	return NewDatabase(c, name)
}

// Transaction 提供了一个在 MongoDB 中执行事务的方法。
// 它接受一个函数作为参数，该函数接受一个 mongo.SessionContext 类型的参数并返回一个 error 类型的值。
// 如果函数返回一个非空的 error，事务将被回滚，否则事务将被提交。
// 如果在启动事务时出现错误，将返回该错误。
// 如果在提交或回滚事务时出现错误，将返回该错误。
func (c *Client) Transaction(fn func(sessionContext mongo.SessionContext) error) error {
	session, err := c.cli.StartSession()
	if err != nil {
		return err
	}
	defer func() {
		session.EndSession(context.Background())
	}()

	var f = func(sessionContext mongo.SessionContext) (interface{}, error) {
		err := fn(sessionContext)
		return nil, err
	}

	_, err = session.WithTransaction(context.Background(), f)
	return err
}
