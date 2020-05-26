/*
This file is under GNU AFFERO GENERAL PUBLIC LICENSE

Permissions of this strongest copyleft license are conditioned
on making available complete source code of licensed works and
modifications, which include larger works using a licensed work,
under the same license. Copyright and license notices must be preserved.
Contributors provide an express grant of patent rights.
When a modified version is used to provide a service over a network,
the complete source code of the modified version must be made available.

Edoardo Ottavianelli, https://edoardoottavianelli.it
*/

package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// TODO
func ConnectDB(connectionString string,databaseName string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(ctx)
	return client
}

// TODO
func InsertUsers(database *mongo.Database, users []User) {
	collection := GetUsers(database)
	usersInt := []interface{}{users}
	insertManyResult, err := collection.InsertMany(context.TODO(), usersInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple users: ", insertManyResult.InsertedIDs)
}

// TODO
func InsertWebsites(database *mongo.Database, websites []Website) {
	collection := GetWebsites(database)
	websitesInt := []interface{}{websites}
	insertManyResult, err := collection.InsertMany(context.TODO(), websitesInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple websites: ", insertManyResult.InsertedIDs)
}

// TODO
func InsertUser(database *mongo.Database, user User) {
	collection := GetUsers(database)
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single user: ", insertResult.InsertedID)
}

// TODO
func InsertWebsite(database *mongo.Database, website Website) {
	collection := GetWebsites(database)
	insertResult, err := collection.InsertOne(context.TODO(), website)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single website: ", insertResult.InsertedID)
}