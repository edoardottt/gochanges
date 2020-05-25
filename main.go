/*
ADD LICENSE
 */

package main

import (
	"github.com/edoardottt/gochanges/scraper"
)

func main() {
	u := "https://www.google.com"
	interval := 3
	m := scraper.AddMonitor(u, interval)
	scraper.StartMonitoring(m)
}