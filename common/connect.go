/*
ADD LICENSE
 */

package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// TODO
type Website struct {

}

func Connect() {
	res, err := http.Get("http://www.google.com/robots.txt")
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
