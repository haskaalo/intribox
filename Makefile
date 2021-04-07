GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)

export AWS_ACCESS_KEY_ID := test
export AWS_SECRET_ACCESS_KEY := test

.PHONY: lint
lint:
	golint $(PACKAGES)

.PHONY: test
test:
	go test -cover github.com/haskaalo/intribox/...
