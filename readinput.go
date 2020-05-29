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
	"bufio"
	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"log"
	"net/mail"
	"net/url"
	"os"
	"errors"
	"strconv"
	"strings"
)

//ReadInput reads from a file
//the emails, websites and their interval time.
//Emails will be used as destinations where sends notifications.
//Website urls have their corresponding interval (integer)
//that is the seconds between two requests.
//Returns also a slice of Monitors, object used for scraping.
func ReadInput(filename string) ([]db.User, []db.Website, []scraper.Monitor) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var users []db.User
	var websites []db.Website
	var monitors []scraper.Monitor

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		address := scanner.Text()
		// if only one field in a row, default time = 5 min
		if len(strings.Fields(address)) == 1 {
			if IsEmail(address) {
				users = append(users, db.User{Email: address[1 : len(address)-1]})
			} else {
				if IsWebsite(address) {
					website := db.Website{
						Address:   address,
						Body:      scraper.GetContent(address),
						Timestamp: scraper.GetCurrentTimestamp(),
					}
					websites = append(websites, website)
					if err != nil {
						log.Fatal(err)
					}
					monitors = append(monitors, scraper.Monitor{
						Website: website,
						Seconds: 300,
					})
				}
			}
		} else if len(strings.Fields(address)) == 2 { //otherwise, check also the second input on the row
			fields := strings.Fields(address)
			if IsWebsite(fields[0]) {
				website := db.Website{
					Address:   address,
					Body:      scraper.GetContent(fields[0]),
					Timestamp: scraper.GetCurrentTimestamp(),
				}
				websites = append(websites, website)
				seconds, err := strconv.Atoi(fields[1])
				if err != nil {
					log.Fatal(err)
				}
				monitors = append(monitors, scraper.Monitor{
					Website: website,
					Seconds: seconds,
				})
			}
		} else {
			err = errors.New("Bad input.")
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return users, websites, monitors
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
