test := go test -v -cover -parallel 4

.PHONY: run test test-short format

run: bin/app
	@bin/app
bin/app: main.go format
	@go build -o bin/app ./main.go
test: main_test.go format
	@$(test)
test-short: main_test.go format
	@$(test) -short
format:
	@goimports -w .
