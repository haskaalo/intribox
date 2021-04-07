GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)

.PHONY: lint
lint:
	golint $(PACKAGES)

.PHONY: test

test: export AWS_ACCESS_KEY_ID = "test"
test: export AWS_SECRET_ACCESS_KEY = "test"
test: export CONFIG_PATH = intribox_config.dev.s3.ini
test:
	go test -cover ./...
