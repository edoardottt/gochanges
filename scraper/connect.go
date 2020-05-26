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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Monitor struct {
	website 	string
	seconds 	int
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
	monitor.seconds = interval
	return monitor
}

// TODO
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
	fmt.Printf("%s", body)	//DEBUG
	return string(body)
}

// TODO
func doEvery(d time.Duration, f func(u string) string, monitor Monitor, change Change) {
	for _ = range time.Tick(d) {
		content := f(monitor.website)
		newTimestamp := GetCurrentTimestamp()
		edited := Edited(change, content)
		if edited {
			// TODO
			fmt.Println("edited")
		}
		fmt.Println(newTimestamp)
		fmt.Println(content)
	}
}

// TODO
func StartMonitoring(monitor Monitor) {
	change := Change{monitor: monitor, timestamp: GetCurrentTimestamp()}
	d := time.Duration(monitor.seconds) * time.Second
	doEvery(d, GetContent, monitor,change)
}