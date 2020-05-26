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
	"github.com/edoardottt/gochanges/db"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Monitor struct {
	Website 	db.Website
	Seconds 	int
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
	body := GetContent(address)
	monitor.Website = db.Website{Address: address, Body: body ,Timestamp: GetCurrentTimestamp()}
	monitor.Seconds = interval
	return monitor
}

// TODO
func AddUser(email string) db.User {
	verified,err := VerifiedEmail(email)
	if !verified {
		log.Fatal(err)
	}
	return db.User{Email: email}
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
func doEvery(d time.Duration, f func(u string) string, monitor Monitor) {
	for _ = range time.Tick(d) {
		content := f(monitor.Website.Address)
		newTimestamp := GetCurrentTimestamp()
		edited := Edited(monitor.Website.Body, content)
		if edited {
			// TODO
			fmt.Println("edited")
			//add website change in mongodb
			//update website struct and add in Monitor
			//email user
		}
		fmt.Println(newTimestamp)
		fmt.Println(content)
	}
}

// TODO
func StartMonitoring(monitor Monitor) {
	d := time.Duration(monitor.Seconds) * time.Second
	doEvery(d, GetContent, monitor)
}

func VerifiedEmail(email string) (bool,error) {
	return true,nil
}