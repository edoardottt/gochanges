SHELL := /bin/bash

#---Run the service---
up:
	docker-compose up

#---Stop the service---
stop:
	docker-compose stop

#---Delete the service---
down:
	docker-compose down

#---Run testcases---
test:
	go test ./...

#---Delete the image created---
clean:
	docker rmi gochanges_gochanges:latest

#---Prune---
prune:
	docker system prune -f

#---Restart the service (applying the changes made)---
restart:
	docker-compose down
	docker rmi gochanges_gochanges:latest
	docker-compose up
