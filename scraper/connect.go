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

// TODO
func ParseUrl(urlPath string)  url.URL {
	u, err := url.Parse(urlPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(*u)			//DEBUG
	return *u
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

// TODO
func HealthCheck(u string) bool {
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	return res.StatusCode == 200
}