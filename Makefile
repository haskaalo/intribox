GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)
export AWS_ACCESS_KEY_ID := "test"
export AWS_SECRET_ACCESS_KEY = "test"
export CONFIG_PATH = $(PWD)/intribox_config.dev.s3.ini
export AWS_DEFAULT_REGION = "us-east-1"

.PHONY: lint
lint:
	golint $(PACKAGES)

.PHONY: test
test:
	go test -cover ./...

.PHONY: setupaws
setupaws: 
	aws --endpoint-url=http://localhost:4566 s3 mb s3://testbucket