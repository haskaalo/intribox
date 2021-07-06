GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)
export AWS_ACCESS_KEY_ID := DevAccessKey
export AWS_SECRET_ACCESS_KEY = DevSecretKey
export AWS_DEFAULT_REGION = us-east-1

.PHONY: lint
lint:
	golint $(PACKAGES)

.PHONY: test
test:
	go test -cover ./...
