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
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/mail"
	"time"
)

// TODO
type Website struct {
	Address 	string
	Body 		string
	Timestamp 	int64
}

// TODO
type User struct {
	Email		string
}


// TODO
func InsertWebsite(client *mongo.Client, website string) (bool, error) {

}

// TODO
func InsertChange(client *mongo.Client, website string) (bool,error) {

}

// TODO
func IsWebsitePresent(client *mongo.Client, website string) (bool,error) {

}