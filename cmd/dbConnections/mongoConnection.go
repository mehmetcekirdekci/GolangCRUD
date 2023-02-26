package dbConnections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

func MongoConnectionOpen() *mongo.Database {
	return mongoConnectionOpen()
}

func mongoConnectionOpen() *mongo.Database {
	var db *mongo.Database
	uri := "mongodb+srv://KozZzy:N5WIQQagC0LJKjP9@cluster0.j9q6b1m.mongodb.net/test"
	databaseName := "Examples"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		println(err)
		return nil
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		println(err)
		return nil
	}
	db = client.Database(databaseName)
	return db
}
