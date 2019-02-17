run := docker-compose run
app := $(run) app
db := $(run) db
make := $(app) make

.PHONY: build test

build:
	@$(make) bin/app
test:
	@$(make) test
go-sh:
	@$(app) bash
db-sh:
	@$(db) bash
