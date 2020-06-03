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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//GetDatabase returns the pointer to the database d(input).
func GetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	database := client.Database(databaseName)

	return database
}

//GetUsers returns the collection of users
func GetUsers(database *mongo.Database) *mongo.Collection {
	return database.Collection("users")
}

//GetWebsites returns the collection of websites
func GetWebsites(database *mongo.Database) *mongo.Collection {
	return database.Collection("websites")
}

//InsertUsers inserts into the collection users
//in database d(input) a slice of users inputted.
func InsertUsers(connString string, dbName string, users []User) {
	client, ctx := ConnectDB(connString)
	database := GetDatabase(client, dbName)
	collection := GetUsers(database)
	usersInt := make([]interface{}, len(users))
	for i := range users {
		usersInt[i] = users[i]
	}
	insertManyResult, err := collection.InsertMany(context.TODO(), usersInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple users: ", insertManyResult.InsertedIDs)
	client.Disconnect(ctx)
}

//InsertWebsites inserts into the collection websites
//in database d(input) a slice of websites inputted.
func InsertWebsites(connString string, dbName string, websites []Website) {
	client, ctx := ConnectDB(connString)
	database := GetDatabase(client, dbName)
	collection := GetWebsites(database)
	websitesInt := make([]interface{}, len(websites))
	for i := range websites {
		websitesInt[i] = websites[i]
	}
	insertManyResult, err := collection.InsertMany(context.TODO(), websitesInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple websites: ", insertManyResult.InsertedIDs)
	client.Disconnect(ctx)
}

//InsertUser inserts an user into the collection users
func InsertUser(connString string, dbName string, user User) {
	client, ctx := ConnectDB(connString)
	database := GetDatabase(client, dbName)
	collection := GetUsers(database)
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single user: ", insertResult.InsertedID)
	client.Disconnect(ctx)
}

//InsertWebsite inserts a website url into the collection websites
func InsertWebsite(connString string, dbName string, website Website) {
	client, ctx := ConnectDB(connString)
	database := GetDatabase(client, dbName)
	collection := GetWebsites(database)
	insertResult, err := collection.InsertOne(context.TODO(), website)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single website: ", insertResult.InsertedID)
	client.Disconnect(ctx)
}

// TODO
func GetAllUsers(connString string, dbName string) []User {
	var users []User
	client, ctx := ConnectDB(connString)
	database := GetDatabase(client, dbName)
	collection := GetUsers(database)
	filterCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = filterCursor.All(ctx, &users)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

// TODO
func GetAllWebsites(connString string, dbName string) []Website {
	var websites []Website
	client, ctx := ConnectDB(connString)
	database := GetDatabase(client, dbName)
	collection := GetWebsites(database)
	filterCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = filterCursor.All(ctx, &websites)
	if err != nil {
		log.Fatal(err)
	}
	return websites
}
