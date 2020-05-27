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
	"strconv"
	"strings"
)

// TODO
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
		if IsEmail(address) && len(strings.Fields(address)) == 1 {
			users = append(users, db.User{Email: address[1 : len(address)-1]})
		} else {
			if IsWebsite(address) && len(strings.Fields(address)) == 2 {
				fields := strings.Fields(address)
				website := db.Website{
					Address:   fields[0],
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
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return users, websites, monitors
}

// TODO
func IsEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

// TODO
func IsWebsite(address string) bool {
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return false
	}
	return true
}
