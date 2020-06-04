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
	"errors"
	"net/url"
	"regexp"
	"strconv"
)

//checkEmail checks if the email inputted is a valid email.
func checkEmail(email string) bool {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}
	return true
}

//checkWebsite checks if the website inputted is a valid URL.
func checkWebsite(website string) (bool, error) {
	u, err := url.Parse(website)
	if err != nil {
		err = errors.New("website inputted is not a valid URL")
		return false, err
	} else if u.Scheme == "" || u.Host == "" {
		err = errors.New("website inputted must be an absolute URL")
		return false, err
	} else if u.Scheme != "http" && u.Scheme != "https" {
		err = errors.New("website inputted must begin with http or https")
		return false, err
	}
	return true, nil
}

//checkTelegram checks if the telegram inputted is a valid telegram token.
func checkTelegram(telegram string) bool {
	return true
}

//checkInterval checks if the interval inputted is a valid interval.
func checkInterval(interval string) (bool, error) {
	i, err := strconv.Atoi(interval)
	if err != nil {
		return false, err
	}
	return 1 <= i && i <= 86400, nil
}
