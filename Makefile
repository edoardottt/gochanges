SHELL := /bin/bash

up:
	docker-compose up

stop:
	docker-compose stop

down:
	docker-compose down

test:
	go test ./...

clean:
	docker rmi gochanges_gochanges:latest

prune:
	docker system prune -f

restart:
	docker-compose down
	docker rmi gochanges_gochanges:latest
	docker-compose up