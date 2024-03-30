package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Parameters that represent a database connection which lasts the entire process.
// This is to be passed around to create instances of transactionConnection
type DatabaseConnection struct {
	ConnectionString string
	DatabaseName     string
}

// A connection to a database that is used by transactions, scoped to each transaction.
type TransactionConnection struct {
	database *mongo.Database
	client   *mongo.Client
	ctx      context.Context
}

// Performs a database transaction, automatically creating the db client and
// cleaning up afterwards.
func (dbc *DatabaseConnection) connect() TransactionConnection {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbc.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(dbc.DatabaseName)

	return TransactionConnection{
		database: database,
		client:   client,
		ctx:      ctx,
	}
}

func (tc *TransactionConnection) cleanup() {
	tc.client.Disconnect(tc.ctx)
}
