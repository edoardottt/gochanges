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
	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"github.com/edoardottt/gochanges/webserver"
	"os"
	"fmt"
	"strconv"
)

func getEnvOrDefault(key string, defaultValue string) string {
	tryValue := os.Getenv(key)
	if len(tryValue) == 0 {
		tryValue = defaultValue
	}
	return tryValue
}

func main() {
	connString := getEnvOrDefault("MONGO_CONN", "mongodb://localhost:27017")
	dbName := getEnvOrDefault("DB_NAME", "gochangesdb")
	port, port_err := strconv.Atoi(getEnvOrDefault("LISTEN_PORT", "3822"))
	if port_err != nil {
		fmt.Println("Please ensure port is an integer! Abort.")
		os.Exit(1)
	}

	// Take all websites into mongodb
	// Start monitoring all websites yet present
	websites := db.GetAllWebsites(connString, dbName)
	for _, website := range websites {
		go scraper.StartMonitoring(website, connString, dbName)
	}
	// Scan input and monitor it
	webserver.StartListen(port)
}
