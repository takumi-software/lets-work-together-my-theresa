SHELL := /bin/bash
PROTOTOOL_VERSION := 1.8.1

proto:
	@echo Formatting proto and generating proto buffer files...
	@docker run --env GOPATH=/go -v $(shell pwd):/go/src/github.com/takumi-software/lets-work-together-my-theresa uber/prototool:$(PROTOTOOL_VERSION) /bin/bash -c 'sleep 5; /go/src/github.com/takumi-software/lets-work-together-my-theresa/proto_generator.sh;'
run: proto
	#It builds the backend application using your current architecture
	docker-compose up
stop:
	docker-compose down
lint: bin/golangci-lint revive
	bin/golangci-lint run -v
.PHONY: lint

revive: bin/revive
	ulimit -n 2048 && bin/revive --config revive.toml -formatter stylish -exclude ./third_party/... github.com/takumi-software/lets-work-together-my-theresa/ ./...
.PHONY: revive

bin/goimports:
	GOBIN=$(CURDIR)/bin go install golang.org/x/tools/cmd/goimports

bin/golangci-lint:
	GOBIN=$(CURDIR)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint

bin/revive:
	GOBIN=$(CURDIR)/bin go get -u github.com/mgechev/revive

bin/prototool:
	GOBIN=$(CURDIR)/bin go install github.com/uber/prototool/cmd/prototool
#	GOBIN=$(CURDIR)/bin go install github.com/bufbuild/buf/cmd/buf

bin/protoc-gen-go:
	GOBIN=$(CURDIR)/bin go install github.com/golang/protobuf/protoc-gen-go
