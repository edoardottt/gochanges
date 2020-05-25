/*
ADD LICENSE
 */

package main

import (
	"github.com/edoardottt/gochanges/scraper"
)

func main() {
	u := "http://www.google.com/robots.txt"
	scraper.ParseUrl(u)
	scraper.GetContent(u)
}