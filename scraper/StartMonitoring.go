package scraper

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/edoardottt/gochanges/db"
)

//Start monitoring a given website.
func StartMonitoring(scrapeTarget db.ScrapeTarget, dbc *db.DatabaseConnection) {
	log.Printf("Starting monitoring of %s", scrapeTarget.Url)

	MonitorAndUpdate(scrapeTarget, dbc)
	for range time.Tick(time.Duration(scrapeTarget.MonitorIntervalSeconds) * time.Second) {
		MonitorAndUpdate(scrapeTarget, dbc)
	}
}

func MonitorAndUpdate(scrapeTarget db.ScrapeTarget, dbc *db.DatabaseConnection) {
	body := GetContent(scrapeTarget.Url)
	newTimestamp := GetCurrentTimestamp()
	edited := checkBodyChanged(scrapeTarget.LastBody, body)
	if edited {
		scrapeTarget.LastBody = body
		scrapeTarget.LastChangedUnixMillis = newTimestamp
	}
	scrapeTarget.LastMonitoredUnixMillis = newTimestamp
	dbc.InsertScrapeTargets([]db.ScrapeTarget{scrapeTarget})
}

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

//GetCurrentTimestamp returns milliseconds
//from 1 Jan 1970
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

//checkBodyChanged tell us if a content of a website is changed
func checkBodyChanged(body string, content string) bool {
	return body != content
}
