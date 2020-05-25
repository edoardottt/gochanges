/*
ADD LICENSE
 */

package scraper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Monitor struct {
	website 	string
	interval 	int
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
func ParseUrl(urlPath string)  (url.URL,error) {
	u, err := url.Parse(urlPath)
	if err != nil {
		return url.URL{},err
	}
	fmt.Println(*u)			//DEBUG
	return *u,nil
}

// TODO
func AddMonitor(address string, interval int) Monitor {
	monitor := Monitor{}
	_, err := ParseUrl(address)
	if err != nil {
		log.Fatal(err)
	}
	monitor.website = address
	monitor.interval = interval
	return monitor
}

// TODO
func GetContent(u string) []byte {
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)	//DEBUG
	return body
}

