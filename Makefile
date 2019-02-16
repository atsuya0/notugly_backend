cmd := docker-compose run app make

.PHONY: build test test-short

build:
	@$(cmd) bin/app
test:
	@$(cmd) test
test-short:
	@$(cmd) test-short
