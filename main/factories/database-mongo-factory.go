package factories

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db_mongo_con *mongo.Client

func NewDatabaseMongoOpenConnection() error {
	uri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	db_mongo_con = client
	return nil
}

func NewCloseDatabaseMongoConnection() error {
	return db_mongo_con.Disconnect(context.Background())
}
