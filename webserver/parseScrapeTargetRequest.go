package webserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type ScrapeTargetRequest struct {
	Url                    string `json:"url"`
	MonitorIntervalSeconds int    `json:"monitorIntervalSeconds"`
}

// Parse a scrape target request
func parseScrapeTargetRequest(r *http.Request, reqBody *ScrapeTargetRequest) error {
	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return errors.New("failed to decode request body")
	}

	// Check if the website is valid
	if err := CheckUrl(reqBody.Url); err != nil {
		return errors.New("invalid url: " + err.Error())
	}

	// Check if the interval is valid
	if err := CheckInterval(reqBody.MonitorIntervalSeconds); err != nil {
		return errors.New("invalid interval: " + err.Error())
	}

	return nil
}

//checkWebsite checks if the website inputted is a valid URL.
func CheckUrl(website string) error {
	u, err := url.Parse(website)
	if err != nil {
		err = errors.New("website inputted is not a valid URL")
		return err
	} else if u.Scheme == "" || u.Host == "" {
		err = errors.New("website inputted must be an absolute URL")
		return err
	} else if u.Scheme != "http" && u.Scheme != "https" {
		err = errors.New("website inputted must begin with http or https")
		return err
	}
	return nil
}

//checkInterval checks if the interval inputted is a valid interval.
func CheckInterval(interval int) error {
	if !(1 <= interval && interval <= 86400) {
		return errors.New("Interval should be between 1s and 1 day")
	}
	return nil
}
