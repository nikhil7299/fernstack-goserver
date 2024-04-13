package database

import (
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var MongoDatabase *mongo.Database
var mongoDBName string

// type MongoInstance struct {
// 	MongoClient   *mongo.Client
// 	MongoDatabase *mongo.Database
// }

func MongoConnect() error {
	mongoUri := os.Getenv("MONGO_URL")
	if mongoUri == "" {

		return errors.New("Set MONGO_URL environment Variable")
	}
	mongoDb := os.Getenv("MONGO_DB")
	if mongoDb == "" {
		return errors.New("Set MONGO_DB environment Variable")
	} else {
		mongoDBName = mongoDb
	}

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		return err
	} else {
		log.Print("Connected to MongoDB")
	}

	MongoDatabase = mongoClient.Database(mongoDBName)

	pingErr := mongoClient.Ping(context.Background(), nil)
	if pingErr != nil {
		return errors.New("Can't verify a connection")
	}
	return nil
}

// func GetCollection(collectionName string) *mongo.Collection {
// 	log.Print(mongoDBName)
// 	log.Print(collectionName)
// 	mongoDatabase := mongoClient.Database(mongoDBName)
// 	log.Print(mongoDatabase.Name())
// 	// return mongoClient.Database(mongoDBName).Collection(collectionName)
// 	return mongoDatabase.Collection(collectionName)
// }

func MongoDisconnect() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}
