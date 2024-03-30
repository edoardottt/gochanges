package db

// A target to be scraped
type ScrapeTarget struct {
	Url                     string `json:"url"`
	LastBody                string `json:"lastBody"`
	MonitorIntervalSeconds  int    `json:"monitorIntervalSeconds"`
	LastMonitoredUnixMillis int64  `json:"lastMonitoredUnixMillis"`
	LastChangedUnixMillis   int64  `json:"lastChangedUnixMillis"`
}
