GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/intribox/... | grep -v /vendor/)
export AWS_ACCESS_KEY_ID := DevAccessKey
export AWS_SECRET_ACCESS_KEY = DevSecretKey
export CONFIG_PATH = $(PWD)/intribox_config.dev.ini
export AWS_DEFAULT_REGION = us-east-1

.PHONY: lint
lint:
	golint $(PACKAGES)

.PHONY: test
test:
	go test -cover ./...

.PHONY: setupaws
setupaws: 
	aws --endpoint-url http://localhost:9000 s3 mb s3://testbucket