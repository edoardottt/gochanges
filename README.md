# gochanges: a website changes tracker for developers

![gochanges logo](images/gochanges.png)

gochanges is a website changes tracker for developers (packaged as a microservice).

## User guide

Requirements: Docker.

```bash
git clone $REPOSITORY

# Start using the standard docker repertoire: up, down, rm etc.
docker-compose up -d

# To watch a new URL:
curl -X POST http://localhost:3822/scrapeTarget \
  -d '{"url":"https://google.com","monitorIntervalSeconds":300}'
# {"result":"success"}

# To see all current scraped targets: (You are probably most interested in the 'lastChanged' field)
curl http://localhost:3822/scrapeTarget
# {"scrapeTargets": [
#   {
#     "url": "https://www.google.com",
#     "lastBody": "...",
#     "monitorIntervalSeconds": 300,
#     "lastMonitoredUnixMillis": 1648719600,
#     "lastChangedUnixMillis": 1648705200
#   }
# ]}
```

## Iterative development

```bash
docker-compose up -d mongodb # required for keeping track of state

go run main.go # edit and rerun this command

# Then you should be able to execute the commands above!

# If you're already listening on another port, or you want to use another mongodb, etc, then check out the environment variables in main.go.
```

## Project Layout

```bash
$ tree -P '*.go' --prune .
.
├── db
│   ├── connection.go               # Wrapper for handling creation of connections 
│   │                               # from db connection string
│   ├── transactions.go             # Business logic transactions
│   └── storedTypes.go              # Domain types to be stored in MongoDB
├── main.go                         # Main file, including configuration parsing
│                                   # and composing together all the functionality.
├── scraper             
│   └── StartMonitoring.go          # Schedules a monitoring task for a given URL.
└── webserver
    ├── parseScrapeTargetRequest.go # Parses and validates a http request to scrape a website.
    └── server.go                   # Maps routes to functionality; starts the server
```

## Credits and project info

Built with <img src="images/docker-logo.png" style="height:1em"/>
<img src="images/golang-logo.png" style="height:1em"/>
<img src="images/mongodb-logo.png" style="height:1em"/>

![Go workflow](https://github.com/edoardottt/gochanges/workflows/Go/badge.svg) [![Go report card](https://goreportcard.com/badge/github.com/edoardottt/gochanges)](https://goreportcard.com/report/github.com/edoardottt/gochanges) ![license AGPLv3.0](images/licenseBadge.svg)

Coded with 💙 by [edoardottt](https://edoardoottavianelli.it), revived by acenturyandabit 2024.

This repository is under [GNU Affero General Public License v3.0](https://github.com/edoardottt/gochanges/blob/master/LICENSE).