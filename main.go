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

package main

import (
	"bufio"
	"fmt"
	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"os"
)

func main() {
	connString := os.Getenv("MONGO_CONN")
	//connString := "mongodb://hostname:27017"

	dbName := os.Getenv("DB_NAME")
	//dbName := "gochangesdb"

	// Take all websites into mongodb
	// Start monitoring all websites yet present
	websites := db.GetAllWebsites(connString, dbName)
	for _, website := range websites {
		go scraper.StartMonitoring(website, connString, dbName)
	}
	// Scan input and monitor it
	fmt.Println("Insert input.")
	fmt.Println("Press CTRL+C to terminate...")
	var input string
	for {
		fmt.Println("-------------------------------------")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter input: ")
		input, _ = reader.ReadString('\n')
		inputtedType, user, website := ReadInput(input)
		//Insert data in mongoDB
		if inputtedType == "user" {
			db.InsertUser(connString, dbName, user)
		} else if inputtedType == "website" {
			db.InsertWebsite(connString, dbName, website)
			go scraper.StartMonitoring(website, connString, dbName)
		} else if inputtedType == "error" {
			fmt.Println("bad input")
		}
		fmt.Println("-------------------------------------")
	}
}
