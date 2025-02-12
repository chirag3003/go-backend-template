package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	session *mongo.Client
}

func (c *conn) Close() {
	if err := c.session.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (c *conn) DB() *mongo.Database {
	return c.session.Database(os.Getenv("MONGO_DB"))
}

func ConnectMongo() Connection {
	session, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connected")
	return &conn{session}
}
