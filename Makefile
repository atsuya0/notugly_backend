run := docker-compose run
make := $(run) app make

.PHONY: build test test-short

build:
	@$(make) bin/app
test:
	@$(make) test
test-short:
	@$(make) test-short
