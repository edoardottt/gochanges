/* This file is under GNU AFFERO GENERAL PUBLIC LICENSE */

package restServer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
)

// HTTP Server, excluding singleton modules (i.e. http and scraper).
type httpServer struct {
	db              db.DatabaseConnection
	scraperInstance scraper.Scraper
	// http
	// scraper
}

func MakeHTTPServer(
	db db.DatabaseConnection,
	scraperInstance scraper.Scraper,
) *httpServer {
	return &httpServer{
		db:              db,
		scraperInstance: scraperInstance,
	}
}

// Add all routes and start the http listener.
func (s *httpServer) StartListen(port int) {
	http.HandleFunc("/scrapeTarget", s.handleScrapeTarget)
	portString := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s", portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}

type scrapeTargetGetResponse struct {
	ScrapeTargets []db.ScrapeTarget `json:"scrapeTargets"`
}

// Handle the /scrapeTarget endpoint, including GET and PUT.
func (s *httpServer) handleScrapeTarget(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a /scrapeTarget request")
	switch r.Method {
	case http.MethodGet:
		scrapeTargets := scrapeTargetGetResponse{s.db.GetScrapeTargets()}
		jsonData, err := json.Marshal(scrapeTargets)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	case http.MethodPost:
		var scrapeTargetRequest ScrapeTargetRequest
		if err := parseScrapeTargetRequest(r, &scrapeTargetRequest); err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		scrapeTarget := db.ScrapeTarget{
			Url:                    scrapeTargetRequest.Url,
			LastBody:               "",
			MonitorIntervalSeconds: scrapeTargetRequest.MonitorIntervalSeconds,
			LastMonitoredUnixSecs:  scraper.GetCurrentTimestamp(),
			LastChangedUnixSecs:    scraper.GetCurrentTimestamp(),
		}
		log.Printf("Got request to monitor new url %s", scrapeTarget.Url)
		go s.scraperInstance.StartMonitoring(scrapeTarget)
		w.Write([]byte("{\"result\":\"success\"}")) // easier to do inline than error check
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
