run := docker-compose run --rm
app := $(run) app
db := $(run) db
make := $(app) make

.PHONY: build test go-sh db-sh

build:
	@$(make) bin/app
test:
	@$(make) test
go-sh:
	@$(app) bash
db-sh:
	@$(db) bash
