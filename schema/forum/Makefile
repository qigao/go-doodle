#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=forum
export LDFLAGS="-w -s"
SHELL   = /bin/bash

all: install-deps gen-mock build test
install-deps:
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
	go install github.com/vektra/mockery/v2@latest
	go install github.com/vburenin/ifacemaker@latest
start-mysql:
	docker-compose -f $(ROOT)/docker/docker-compose-schema.yml up -d

gen-schema:
	go build -o bin/dbimport $(ROOT)/cmd/dbimport
	sleep 5
	bin/dbimport -db $(ROOT)/sql -path $(ROOT)

gen-orm:
	@sqlboiler mysql

stop-mysql:
	@docker-compose -f $(ROOT)/docker/docker-compose-schema.yml down

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
