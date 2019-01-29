app := notugly
source := ./app
credentials := $(source)/credentials.json

.PHONY: run build up down

run: build
	@FIREBASE_AUTH_CREDENTIALS=$(credentials) ./$(app)
build:
	@goimports -w $(source) ./main.go
	@go build -o $(app)
up:
	@docker-compose up
down:
	@docker-compose down
