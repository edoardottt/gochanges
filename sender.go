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
	"log"
	"net/smtp"
)

func SendEmail(emails []*db.User) {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "piotr@mailtrap.io", "extremely_secret_pass", "smtp.mailtrap.io")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"billy@microsoft.com"}
	msg := []byte("To: billy@microsoft.com\r\n" +
		"Subject: Why are you not using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here’s the space for our great sales pitch\r\n")
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, "piotr@mailtrap.io", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}