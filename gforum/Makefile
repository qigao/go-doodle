.PHONY: buf-build buf-lint buf-generate unittest e2etest
buf-build: # compile check
	buf build ./grpc/proto --error-format=json

.PHONY: buf-lint
buf-lint: # lint Protobuf files
	buf lint ./grpc/proto --error-format=json

.PHONY: buf-generate
buf-generate: # generate Go code
	cd grpc &&  buf generate --error-format=json

unittest:
	go test -v ./handler -parallel 4

e2etest:
	bash test/run-api-tests.sh
