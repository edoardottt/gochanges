package scraper

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/edoardottt/gochanges/db"
)

type Scraper struct {
	onScrapeFunctions [](func(db.ScrapeTarget))
}

func BuildScraper(onScrapeFunctions [](func(db.ScrapeTarget))) Scraper {
	return Scraper{
		onScrapeFunctions: onScrapeFunctions,
	}
}

//Start monitoring a given website.
func (s *Scraper) StartMonitoring(scrapeTarget db.ScrapeTarget) {
	log.Printf("Starting monitoring of %s", scrapeTarget.Url)

	s.MonitorAndUpdate(scrapeTarget)
	for range time.Tick(time.Duration(scrapeTarget.MonitorIntervalSeconds) * time.Second) {
		s.MonitorAndUpdate(scrapeTarget)
	}
}

func (s *Scraper) MonitorAndUpdate(scrapeTarget db.ScrapeTarget) {
	body := GetContent(scrapeTarget.Url)
	newTimestamp := GetCurrentTimestamp()
	edited := checkBodyChanged(scrapeTarget.LastBody, body)
	scrapeTarget.LastMonitoredUnixMillis = newTimestamp
	if edited {
		scrapeTarget.LastBody = body
		scrapeTarget.LastChangedUnixMillis = newTimestamp
	}
	for _, fn := range s.onScrapeFunctions {
		fn(scrapeTarget)
	}
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
