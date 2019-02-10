.PHONY: up down build test format

up:
	@docker info &> /dev/null || sudo systemctl start docker
	@docker-compose up
down:
	@docker-compose down
build: format
	@docker info &> /dev/null || sudo systemctl start docker
	@docker-compose run app make build
test: format
	@docker info &> /dev/null || sudo systemctl start docker
	@docker-compose run app go test -v
format:
	@goimports -w ./app
