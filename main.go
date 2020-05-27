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
	"fmt"
	//"github.com/edoardottt/gochanges/scraper"
	"github.com/edoardottt/gochanges/db"
)

func main() {
	fileName := "example/example1.txt"
	connString := "mongodb://127.0.0.1:27017"
	dbName := "gochangesdb"
	users,websites, _ := ReadInput(fileName)
/*
	fmt.Println(users)
	fmt.Println(websites)
	fmt.Println(monitors)
*/
	//INSERT ALL DATA IN MONGO
	db.InsertUsers(connString,dbName,users)
	db.InsertWebsites(connString,dbName,websites)
	emails := db.GetAllEmails(users)

	fmt.Println(emails)
	//for _,monitor := range monitors {
		//go scraper.StartMonitoring()
	//}

	var e int
	fmt.Scanf("%d", &e)
}
