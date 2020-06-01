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

package main

import (
	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"log"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
)

//ReadInput reads from a file
//the emails, websites and their interval time.
//Emails will be used as destinations where sends notifications.
//Website urls have their corresponding interval (integer)
//that is the seconds between two requests.
//Returns also a slice of Monitors, object used for scraping.
func ReadInput(address string) (string, db.User, db.Website) {

	// if only one field in a row, default time = 5 min
	if len(strings.Fields(address)) == 1 {
		if IsEmail(address) {
			return "user",db.User{Email: address[1 : len(address)-1]},db.Website{}

		} else {
			if IsWebsite(address) {
				website := db.Website{
					Address:   	address,
					Body:      	scraper.GetContent(address),
					Seconds:	300,
					Timestamp: 	scraper.GetCurrentTimestamp(),
				}
				return "website",db.User{},website
			}
		}
	} else if len(strings.Fields(address)) == 2 { //otherwise, check also the second input on the row
		fields := strings.Fields(address)
		if IsWebsite(fields[0]) {
			seconds, err := strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}
			website := db.Website{
				Address:   	address,
				Body:      	scraper.GetContent(fields[0]),
				Seconds:	seconds,
				Timestamp: 	scraper.GetCurrentTimestamp(),
			}
			return "website",db.User{},website
		}
	}
	return "error",db.User{},db.Website{}
}

//IsEmail tell us if a string is an email or not.
//format: <user@mail.com>
func IsEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

//IsWebsite tell us if a string is a website url or not.
//format: https://www.edoardoottavianelli.it/
func IsWebsite(address string) bool {
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return false
	}
	return true
}
