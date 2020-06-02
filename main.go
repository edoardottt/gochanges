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
	"fmt"
	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"io/ioutil"
	"log"
	"net/http"
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
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/save/", handlerSave)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO
func handlerHome(w http.ResponseWriter, r *http.Request) {
	page, _ := loadPage("fe/home.html")
	fmt.Fprintf(w, "%s", page)
}

// TODO
func handlerSave(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	telegram := r.FormValue("telegram")
	website := r.FormValue("website")
	interval := r.FormValue("interval")
	fmt.Println(email)
	fmt.Println(telegram)
	fmt.Println(website)
	fmt.Println(interval)
	fmt.Fprintf(w, "%s %s %s %s", email,telegram, website, interval)
}

// TODO
func loadPage(filename string) (string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}