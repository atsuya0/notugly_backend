go := docker-compose run app go
test := $(go) test -v -cover -parallel 4

.PHONY: up down build test format docker-start

up: docker-start
	@docker-compose up
build: docker-start
	@docker-compose up --build
down:
	@docker-compose down
go-build: docker-start format
	@$(go) build -o app
test: docker-start format
	@$(test)
test-short: docker-start format
	@$(test) -short
format:
	@goimports -w ./src
docker-start:
	@docker info &> /dev/null || sudo systemctl start docker
