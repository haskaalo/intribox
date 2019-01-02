GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)

lint:
	golint $(PACKAGES)

.PHONY: test
test:
	go test -cover github.com/haskaalo/intribox/...
