#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=forum
export LDFLAGS="-w -s"

all: gen-mock build test
run-mysql:
	cp .env $(ROOT)/docker/schema.env
	cp ../shared/containers/db/my.cnf $(ROOT)/docker/my.cnf
	docker-compose -f $(ROOT)/docker/docker-compose-schema.yml up -d
build-immigrate:
	 @go build $(ROOT)/cmd/immigrateDb
immigrate-db: run-mysql build-immigrate
	chmod +x $(ROOT)/docker/wait-for-it.sh
	$(ROOT)/docker/wait-for-it.sh localhost:3306
	./immigrateDb -db $(ROOT)/db/sql -path $(ROOT)
gen-schema: immigrate-db
	sqlboiler mysql
stop-mysql:
	docker-compose -f $(ROOT)/docker/docker-compose-schema.yml down
	rm -rf $(ROOT)/docker/my.cnf
	rm -rf $(ROOT)/docker/schema.env
gen-mock:
	mockery --config test/mockery.yaml

build:
	go build -race  .

build-static:
	CGO_ENABLED=1 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

############################################################
# Test
############################################################

test:
	go test -v -race ./...

container:
	cp -R ../shared .
	docker build  -t echo-realworld . -f docker/Dockerfile
	rm -R shared

run-container:
	docker run --rm -it echo-realworld

.PHONY: build run build-static test container
