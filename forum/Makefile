#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=forum
export LDFLAGS="-w -s"
SHELL   = /bin/bash

all:  gen-mock build test
start-mysql:
	cp .env $(ROOT)/docker/schema.env
	cp ../shared/containers/db/my.cnf $(ROOT)/docker/my.cnf
	docker-compose -f $(ROOT)/docker/docker-compose-schema.yml up -d

gen-schema:
	@go build -o bin/immigrateDb $(ROOT)/cmd/immigrateDb
	sleep 5
	bin/immigrateDb -db $(ROOT)/db/sql -path $(ROOT)

gen-orm:
	@sqlboiler mysql

stop-mysql:
	@docker-compose -f $(ROOT)/docker/docker-compose-schema.yml down
	rm -rf $(ROOT)/docker/my.cnf
	rm -rf $(ROOT)/docker/schema.env

gen-mock:
	@mkdir  -vp mock/service
	ifacemaker -f service/article/service_article.go -s ServiceArticle -i ServiceArticle -p service -c "DONT EDIT: Auto generated" -o service/article.go
	ifacemaker -f service/user/service_user.go -s ServiceUser -i ServiceUser -p service -c "DONT EDIT: Auto generated" -o service/user.go
	mockery --config repository/mockery.yaml
	mockery --config service/mockery.yaml

build:
	go build -o bin/forum -race  .

build-static:
	CGO_ENABLED=1 go build -race -v -o bin/$(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

############################################################
# Test
############################################################

test:
	go test -v -race ./...

container:
	cp -R ../shared .
	docker build  -t echo-forum . -f docker/Dockerfile
	rm -R shared

run-container:
	docker run --rm -it echo-forum

.PHONY: build run build-static test container
