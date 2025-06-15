package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database *mongo.Database
)

const (
	mongoName = "aia"
	mongoUrl  = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.5.2"
)

func Init(ctx context.Context) *mongo.Database {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	Database = client.Database(mongoName)
	return Database
}

func GetCollection(colName string) *mongo.Collection {
	return Database.Collection(colName)
}
