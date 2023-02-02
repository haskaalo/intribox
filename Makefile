GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)
export AWS_ACCESS_KEY_ID := anything_test
export AWS_SECRET_ACCESS_KEY = anything_test
export CONFIG_PATH = $(PWD)/intribox_config.dev.ini
export AWS_DEFAULT_REGION = us-east-1

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -cover ./...

.PHONY: setupaws
setupaws: 
	aws --endpoint-url=http://localhost:4566 s3 mb s3://testbucket