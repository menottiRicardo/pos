package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database settings (insert your own database name and connection URI)
const dbName = "pos"
const mongoURI = "mongodb://root:example@localhost:27017/"

func ConnectWithMongodb() *mongo.Client {

	// load .env file
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }
	// // set mongodb connection string
	// uri := os.Getenv("MONGO_URI")
	// if uri == "" {
	// 	log.Fatal("MONGODB_URI is not set")
	// }
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to mongodb")
	}

	return client
}

var Client *mongo.Client = ConnectWithMongodb()

// OpenCollection get collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("candy").Collection(collectionName)
	return collection
}
