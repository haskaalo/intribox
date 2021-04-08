GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)

.PHONY: lint
lint:
	golint $(PACKAGES)

.PHONY: test
test: export AWS_ACCESS_KEY_ID = "test"
test: export AWS_SECRET_ACCESS_KEY = "test"
test: export CONFIG_PATH = $(PWD)/intribox_config.dev.s3.ini
test:
	go test -cover ./...

.PHONY: setupaws
setupaws: 
	aws --endpoint-url=http://localhost:4566 s3 mb s3://testbucket