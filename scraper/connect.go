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
	return string(body)
}

// TODO
func doEvery(d time.Duration, f func(u string) string, monitor Monitor,emails []string, connString string, dbName string) {
	for _ = range time.Tick(d) {
		content := f(monitor.Website.Address)
		newTimestamp := GetCurrentTimestamp()
		edited := Edited(monitor.Website.Body, content)
		if edited {
			fmt.Println("edited")
			monitor.Website.Body = content
			monitor.Website.Timestamp = newTimestamp
			db.InsertWebsite(connString,dbName, monitor.Website)
			SendEmail(emails)
		}
	}
}

// TODO
func StartMonitoring(monitor Monitor,emails []string, connString string, dbName string) {
	d := time.Duration(monitor.Seconds) * time.Second
	doEvery(d, GetContent, monitor,emails, connString, dbName)
}

