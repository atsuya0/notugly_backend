app := notugly
source := ./app
credentials := $(source)/credentials.json

.PHONY: run build docker-up docker-down

run: build
	@FIREBASE_AUTH_CREDENTIALS=$(credentials) ./$(app)
build:
	@goimports -w $(source) ./main.go
	@go build -o $(app)
docker-up:
	@docker info &> /dev/null || sudo systemctl start docker
	@docker-compose up
docker-down:
	@docker-compose down
