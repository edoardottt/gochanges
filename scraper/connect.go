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
	"github.com/edoardottt/gochanges/db"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//GetContent returns the content of a website
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

//doEvery checks regularly every n seconds (input)
//the content of the website w.
//If the content is changed:
//1. Update the website struct locally
//2. Insert the website changed in mongoDB
//3. Notify all the users
func doEvery(d time.Duration, f func(u string) string, website db.Website, connString string, dbName string) {
	for _ = range time.Tick(d) {
		content := f(website.Address)
		newTimestamp := GetCurrentTimestamp()
		edited := Edited(website.Body, content)
		if edited {
			website.Body = content
			website.Timestamp = newTimestamp
			db.InsertWebsite(connString, dbName, website)
			SendEmail()
		}
	}
}

//StartMonitoring prepare the interval between
//two requests and start monitoring the website w(input).
func StartMonitoring(website db.Website, connString string, dbName string) {
	d := time.Duration(website.Seconds) * time.Second
	doEvery(d, GetContent, website, connString, dbName)
}
