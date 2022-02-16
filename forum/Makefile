#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=forum
export LDFLAGS="-w -s"

all: gen-db gen-mock build test

gen-db:
	cd $(ROOT)/db && sqlboiler mysql

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
