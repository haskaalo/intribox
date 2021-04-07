GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)

export AWS_ACCESS_KEY_ID := test
export AWS_SECRET_ACCESS_KEY := test

.PHONY: lint
lint:
	golint $(PACKAGES)

.PHONY: test
test:
	aws --endpoint-url=http://localhost:4566 s3 mb s3://testbucket
	go test -cover github.com/haskaalo/intribox/...
