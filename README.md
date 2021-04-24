<p align="center">
  <b>
  ⚠️ARCHIVED (AT LEAST FOR NOW). NO LONGER MANTAINED.⚠️
  </b>
</p>


<p align="center">
  <!-- logo -->
  <img src="https://github.com/edoardottt/gochanges/blob/master/images/gochanges.png">
  <b>Fast, scalable, easy to use website changes tracker</b><br>
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
    <!-- mainteinance -->
  <a href="https://edoardoottavianelli.it">
    <img src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" alt="Mainteinance" />
  </a>
      <!-- open-issues -->
  <a href="https://edoardoottavianelli.it">
    <img src="https://img.shields.io/github/issues/Naereen/StrapDown.js.svg" alt="open issues" />
  </a>
  
  <br>
  
  <!-- pr-welcome -->
  <a href="https://edoardoottavianelli.it">
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="pr-welcome" />
  </a>
  <!-- ask-me-anything -->
  <a href="https://edoardoottavianelli.it">
    <img src="https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg" alt="ask me anything" />
  </a>
  <!-- license AGPLv3.0 -->
  <a href="https://github.com/edoardottt/gochanges/blob/master/LICENSE">
    <img src="https://github.com/edoardottt/gochanges/blob/master/images/licenseBadge.svg" alt="License" />
  </a>

  <br>
  <!-- Say thanks -->
  <a href="https://saythanks.io/to/edoardott%40gmail.com">
    <img src="https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg" alt="say thanks" />
  </a>
    <!-- Open-source-love -->
  <a href="https://edoardoottavianelli.it">
    <img src="https://badges.frapsoft.com/os/v2/open-source.png?v=103" alt="opensourcelove" />
  </a>
  <!-- Tweet -->
  <a href="https://twitter.com/intent/tweet?text=Try%20this%20amazing%20tool!%20Just%20love%20it!%20https%3A%2F%2Fgithub.com%2Fedoardottt%2Fgochanges">
    <img src="https://img.shields.io/twitter/url/http/shields.io.svg?style=social" alt="tweet" />
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

Requirements 🔍
----------

- docker
- docker-compose

**Tested on my lightweight laptop, THE ENGINE correctly monitors over 150 websites, the backend is not communicating with the frontend.**

Use these:

- https://github.com/n1trux/awesome-sysadmin#monitoring
- https://github.com/statping/statping
- https://github.com/topics/website-monitor
- https://github.com/brotandgames/ciao
- https://github.com/710leo/urlooker
- https://github.com/ParryQiu/GuGuJianKong
- https://github.com/JuanmaMenendez/website-change-monitor


Get Started 🎉
-------

1. Download this repo
2. Open a terminal and type `docker info && sudo ip addr show docker0`. Find the url of the network.
  You can see mine works on `172.17.0.1`
3. Based on you connection url, edit the environment variable `MONGO_CONN` (Dockerfile file)
4. Type in your terminal `make up`
5. Enjoy!

Common errors here : [COMMON ERRORS](https://github.com/edoardottt/gochanges/blob/master/COMMON_ERRORS.md)

Contributing 🛠
-------

Just open an issue/pull request. See also [CONTRIBUTING.md](https://github.com/edoardottt/gochanges/blob/master/CONTRIBUTING.md) and [CODE OF CONDUCT.md](https://github.com/edoardottt/gochanges/blob/master/CODE_OF_CONDUCT.md)

**To Do:**

- [ ] Understandable frontend
- [ ] Move to hmtl template
- [ ] Data in frontend (emails, telegram tokens and websites tracked)
- [ ] Dynamic Home Page
- [ ] Tests
- [ ] Integrate Telegram-botAPI
- [ ] Send mail using OAuth2.0
- [x] Add comments
- [x] Storing data into MongoDB
- [x] Dockerfile
- [x] docker-compose.yml
- [x] Default interval (5 min)
- [x] Common errors
- [x] Mount external volume
- [x] Read input by localhost:xxxx

License 📝
-------

This repository is under [GNU Affero General Public License v3.0](https://github.com/edoardottt/gochanges/blob/master/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) for contact me.
