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
	body		[]byte
	timestamp	int64
}

// TODO
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

