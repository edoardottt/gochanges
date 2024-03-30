package db

// A target to be scraped
type ScrapeTarget struct {
	Url                    string `json:"url"`
	LastBody               string `json:"lastBody"`
	MonitorIntervalSeconds int    `json:"monitorIntervalSeconds"`
	LastMonitoredUnixSecs  int64  `json:"lastMonitoredUnixSecs"`
	LastChangedUnixSecs    int64  `json:"lastChangedUnixSecs"`
}
