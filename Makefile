run := docker-compose run
make := $(run) app make

.PHONY: build test

build:
	@$(make) bin/app
test:
	@$(make) test
