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

package webserver

import (
	"fmt"
	"github.com/edoardottt/gochanges/db"
	"github.com/edoardottt/gochanges/scraper"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// TODO
func StartListen() {
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/save/", handlerSave)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO
func handlerHome(w http.ResponseWriter, r *http.Request) {

	setContentType(w, r)

	URI := r.RequestURI
	if URI == "/" {
		URI = "./fe/home.html"
	} else {
		URI = "." + URI
	}

	page, _ := loadPage(URI)

	fmt.Fprintf(w, "%s", page)
}

// TODO
func handlerSave(w http.ResponseWriter, r *http.Request) {
	connString := os.Getenv("MONGO_CONN")
	//connString := "mongodb://hostname:27017"

	dbName := os.Getenv("DB_NAME")
	//dbName := "gochangesdb"

	email := r.FormValue("email")
	telegram := r.FormValue("telegram")
	website := r.FormValue("website")
	interval := r.FormValue("interval")

	// CHECK INPUT
	emailOk := CheckEmail(email)
	telegramOk := checkTelegram(telegram)
	websiteOk,err := CheckWebsite(website)
	intervalOk,err := CheckInterval(interval)

	// IF EMAIL OK, INSERT EMAIL
	if emailOk {
		user := db.User{Email: email}
		db.InsertUser(connString, dbName, user)
	}

	// IF TELEGRAM OK, INSERT TELEGRAM
	if telegramOk {
		fmt.Println(telegram)
	}

	if websiteOk && intervalOk {
		sec,_ := strconv.Atoi(interval)
		websiteS := db.Website{Address: website, Body: scraper.GetContent(website), Timestamp: scraper.GetCurrentTimestamp(), Seconds: sec}
		db.InsertWebsite(connString, dbName, websiteS)
	}

	Websites := db.GetAllWebsites(connString, dbName)
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"500 | Internal Server Error")
		return
	}
	// only if there are some data available print on page
	if len(Websites) != 0 {
		tmpl.Execute(w, Websites)
	}

	// DEBUG PRINTING
	fmt.Println("Email:", email)
	fmt.Println("Telegram:", telegram)
	fmt.Println("Website:", website)
	fmt.Println("Interval:", interval)
	fmt.Fprintf(w, "%s %s %s %s", email, telegram, website, interval)
}

// TODO
func loadPage(filename string) (string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func setContentType(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	contentType := "text/html"

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	}

	w.Header().Set("Content-Type", contentType)

}
