/*
ADD LICENSE
 */

package scraper

import (
	"time"
)

// TODO
type Change struct {
	monitor 	Monitor
	body		string
	timestamp	int64
}

// TODO
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// TODO
func Edited(change Change, content string) bool {
	return change.body != content
}