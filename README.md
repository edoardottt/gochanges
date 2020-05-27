
<p align="center">
  <!-- logo -->
  <img src="https://github.com/edoardottt/gochanges/blob/master/images/gochanges.png">
  <b>Fast, scalable website changes tracker</b><br>
  <sub>
    Coded with 💙 by edoardottt.
  </sub>
</p>
<!-- badges -->
<p align="center">
  <!-- build -->
  <a href="#">
    <img src="https://github.com/edoardottt/gochanges/workflows/Go/badge.svg" alt="Go workflows" />
  </a>
  <!-- go report card -->
  <a href="https://goreportcard.com/report/github.com/edoardottt/gochanges">
    <img src="https://goreportcard.com/badge/github.com/edoardottt/gochanges" alt="Go report card" />
  </a>
  <!-- license AGPLv3.0 -->
  <a href="https://github.com/edoardottt/gochanges/blob/master/LICENSE">
    <img src="https://github.com/edoardottt/gochanges/blob/master/images/licenseBadge.svg" alt="License" />
  </a>

  <br>
  Built with<br>
  <!-- docker logo-->
  <a href="https://docker.com">
    <img widht="48" height="48" src="https://github.com/edoardottt/gochanges/blob/master/images/docker-logo.png" alt="docker-logo" />
  </a>
  <!-- go logo-->
  <a href="https://golang.org">
    <img widht="48" height="48" src="https://github.com/edoardottt/gochanges/blob/master/images/golang-logo.png" alt="go-logo" />
  </a>
  <!-- mongodb logo-->
  <a href="https://mongodb.com">
    <img widht="48" height="48" src="https://github.com/edoardottt/gochanges/blob/master/images/mongodb-logo.png" alt="mongodb-logo" />
  </a>
</p>

Requirements
----------

docker, docker-compose

Usage
-------

1. Download this repo
2. Open a terminal and type `sudo ip addr show docker0`. Find the url of the network.
  You can see mine works on 172.17.0.1
3. Based on you connection url, edit the environment variable MONGO_CONN (Dockerfile file)
4. Create (or edit example/example1.txt file) a file that tells the app the emails (receivers) 
  and which websites have to monitor and edit the env. var. FILE_NAME 
6. Type in your terminal `docker-compose up`
7. Enjoy!


**Tested on my lightweight laptop, correctly monitors over 500 websites**
<hr>

**To Do:**

- [x] Add comments
- [x] Storing data into MongoDB
- [x] Dockerfile
- [x] docker-compose.yml
- [ ] Tests
- [ ] Mount ext volume
- [ ] Use only a client to store all changes 
- [ ] Send mail using OAuth2.0
