# gochanges: a website changes tracker for developers

![gochanges logo](images/gochanges.png)

gochanges is a website changes tracker for developers (packaged as a microservice). It is broken up into a backend with a clean API and a frontend.

## User guide

Requirements: Docker.

```bash
git clone $REPOSITORY

docker-compose up -d

# Then go to http://localhost:3822 to see the UI; or you can talk to the server using http:

# To watch a new URL:
curl -X POST http://localhost:3822/add-watch \
  -d '{"url":"https://google.com","intervalMs":60000}'
# {"result":"success"}

# To see what has changed:
curl http://localhost:3822/tracked-sites 
# {"tracked-sites": [
#   "https://google.com": {"lastChanged": "2024-03-29T07:41:51.296Z"} 
# ]}

# To watch a new URL and send an email when it has changed:
curl -X POST http://localhost:3822/add-watch \
  -d '{"url":"https://google.com","mailTo":"user.name@provider.com"}'
# {"result":"success"}
```

## Credits and project info

Built with <img src="images/docker-logo.png" style="height:1em"/>
<img src="images/golang-logo.png" style="height:1em"/>
<img src="images/mongodb-logo.png" style="height:1em"/>

![Go workflow](https://github.com/edoardottt/gochanges/workflows/Go/badge.svg) [![Go report card](https://goreportcard.com/badge/github.com/edoardottt/gochanges)](https://goreportcard.com/report/github.com/edoardottt/gochanges) ![license AGPLv3.0](images/licenseBadge.svg)

Coded with 💙 by [edoardottt](https://edoardoottavianelli.it), revived by acenturyandabit 2024.

This repository is under [GNU Affero General Public License v3.0](https://github.com/edoardottt/gochanges/blob/master/LICENSE).