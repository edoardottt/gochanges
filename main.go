package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"github.com/edoardottt/gochanges/webserver"
)

type config struct {
	DatabaseConnection db.DatabaseConnection
	Port               int
}

func getEnvOrDefault(key string, defaultValue string) string {
	tryValue := os.Getenv(key)
	if len(tryValue) == 0 {
		tryValue = defaultValue
	}
	return tryValue
}

func main() {
	config := getConfigFromEnvironment()

	// Take all websites into mongodb
	// Start monitoring all websites yet present
	websites := config.DatabaseConnection.GetScrapeTargets()
	for _, website := range websites {
		go scraper.StartMonitoring(website, &config.DatabaseConnection)
	}

	httpServer := webserver.MakeHTTPServer(config.DatabaseConnection)
	httpServer.StartListen(config.Port)
}

func getConfigFromEnvironment() config {
	connectionString := getEnvOrDefault("MONGO_CONN", "mongodb://localhost:27017")
	databaseName := getEnvOrDefault("DB_NAME", "gochangesdb")
	Port, port_err := strconv.Atoi(getEnvOrDefault("LISTEN_PORT", "3822"))
	if port_err != nil {
		fmt.Println("Please ensure port is an integer! Abort.")
		os.Exit(1)
	}
	return config{
		db.DatabaseConnection{
			ConnectionString: connectionString,
			DatabaseName:     databaseName,
		},
		Port,
	}
}
