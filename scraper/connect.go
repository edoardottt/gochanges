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

package scraper

import (
	"fmt"
	"github.com/edoardottt/gochanges/db"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Monitor struct {
	Website 	db.Website
	Seconds 	int
}

// TODO
func (Monitor) HealthCheck(u string) bool {
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	return res.StatusCode == 200
}

// TODO
func CreateMonitor(website db.Website, interval int) Monitor {
	return Monitor{Website: website, Seconds: interval}
}

// TODO
func GetContent(u string) string {
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)	//DEBUG
	return string(body)
}

// TODO
func doEvery(d time.Duration, database *mongo.Database, f func(u string) string, monitor Monitor,emails []*db.User) {
	for _ = range time.Tick(d) {
		content := f(monitor.Website.Address)
		newTimestamp := GetCurrentTimestamp()
		edited := Edited(monitor.Website.Body, content)
		if edited {
			fmt.Println("edited")
			monitor.Website.Body = content
			monitor.Website.Timestamp = newTimestamp
			db.InsertWebsite(database, monitor.Website)
			SendEmail(emails)
		}
	}
}

// TODO
func StartMonitoring(monitor Monitor, database *mongo.Database) {
	d := time.Duration(monitor.Seconds) * time.Second
	emails := db.GetAllEmails(database)
	doEvery(d, database, GetContent, monitor,emails)
}