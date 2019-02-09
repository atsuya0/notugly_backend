app := notugly
source := ./app
credentials := $(source)/credentials.json

.PHONY: run build test format docker-up docker-down

run: build
	@FIREBASE_AUTH_CREDENTIALS=$(credentials) ./$(app)
build: format
	@go build -o $(app)
test: format
	@go test -v
format:
	@goimports -w $(source) ./main.go ./main_test.go
docker-up:
	@docker info &> /dev/null || sudo systemctl start docker
	@docker-compose up
docker-down:
	@docker-compose down
