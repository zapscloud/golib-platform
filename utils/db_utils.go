package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

//OpenMongoDbConnection get connection of mongodb
func OpenMongoDbConnection() (*mongo.Client, error) {
	var err error

	log.Println("OpenMongoDbConnection :: Begin")

	dbuser := os.Getenv("MONGO_DB_USER")
	dbsecret := os.Getenv("MONGO_DB_SECRET")
	dbserver := os.Getenv("MONGO_DB_SERVER")

	dburl := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", dbuser, dbsecret, dbserver)
	client, err := mongo.NewClient(options.Client().ApplyURI(dburl))
	if err != nil {
		log.Println("OpenMongoDbConnection :: Client Creation Error")
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println("OpenMongoDbConnection :: Connection Check Failed")
		log.Fatal(err)
	}

	log.Println("OpenMongoDbConnection :: End")

	return client, nil
}

// GetMongoDbCollection - Collection return
func GetMongoDbCollection(client *mongo.Client, collectionName string) (*mongo.Collection, error) {

	dbname := os.Getenv("MONGO_DB_NAME")
	log.Printf("GetMongoDbCollection :: Begin %s %s", dbname, collectionName)

	collection := client.Database(dbname).Collection(collectionName)

	log.Println("GetMongoDbCollection :: end")

	return collection, nil
}

// CloseMongoDb - Close db connection
func CloseMongoDb(client *mongo.Client) error {

	log.Println("CloseMongoDb :: Begin")
	if client == nil {
		log.Println("Connection to MongoDB not open.")
		log.Println("CloseMongoDb :: End")
		return nil
	}
	// Close the connection once no longer needed
	err := client.Disconnect(context.Background())

	if err != nil {
		log.Println("CloseMongoDb :: Disconnect Error")
		log.Fatal(err)
	} else {
		log.Println("Connection to MongoDB closed.")
	}

	log.Println("CloseMongoDb :: End")
	return nil
}
