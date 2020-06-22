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

//Website
type Website struct {
	Address   string `json:"address Str"`
	Body      string `json:"body Str"`
	Seconds   int    `json:"seconds Int"`
	Timestamp int64  `json:"timestamp Int"`
}

//User
type User struct {
	Email string `json:"email Str"`
}

//ConnectDB creates and returns a client connected by a
//connection string to mongoDB.
//Also checks the connection if everything is ok.
func ConnectDB(connectionString string) (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!") //DEBUG
	return client, ctx
}
