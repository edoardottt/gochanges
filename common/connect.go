/*
ADD LICENSE
 */

package common

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

	fmt.Println(u)
	return *u
}

// TODO
func Connect(u string) {

	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
