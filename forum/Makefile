#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=forum
export LDFLAGS="-w -s"
SHELL   = /bin/bash

all: install-deps gen-mock build test
install-deps:
	go install github.com/vburenin/ifacemaker@latest
gen-mock:
	@mkdir  -vp mock/service
		ifacemaker -f repository/mysql/article.go -s ArticleRepo -i IRepoArticle -p repository  -o repository/article.go
		ifacemaker -f repository/mysql/user.go -s UserRepo -i IRepoUser -p repository -o repository/user.go
		ifacemaker -f service/article/service_article.go -s Service -i IServiceArticle -p service  -o service/article.go
		ifacemaker -f service/user/service_user.go -s Service -i IServiceUser -p service  -o service/user.go
		mockery --config repository/mockery.yaml
		mockery --config service/mockery.yaml

build:
	go build -o bin/forum -race  .

build-static:
	CGO_ENABLED=1 go build -race -v -o bin/$(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

test:
	go test -v -race ./...

fmt:
	go fmt  ./...

container:
	cp -R ../shared .
	docker build  -t echo-forum . -f docker/Dockerfile
	rm -R shared

run-container:
	docker run --rm -it echo-forum

.PHONY: build run build-static test container
