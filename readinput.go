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
	"fmt"
	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"log"
	"net/mail"
	"net/url"
	"os"
)

// TODO
func ReadInput(filename string) ([]db.User,[]db.Website) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var users []db.User{}
	var websites []db.Website{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		address := scanner.Text()
		if IsEmail(address) {
			users = append(users,db.User{Email: address})
		}else {
			if IsWebsite(address) {
				websites = append(websites,db.Website{
					Address: address,
					Body: scraper.GetContent(address),
					Timestamp: scraper.GetCurrentTimestamp(),
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return users,websites
}

// TODO
func IsEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

// TODO
func IsWebsite(address string) bool {
	_, err := url.ParseRequestURI(address)
	if err != nil {
		log.Fatal(err)
	}
	return true
}