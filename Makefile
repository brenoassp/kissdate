#!make
GOPATH=$(shell go env GOPATH)
PATH=$(GOPATH)/bin:$(shell echo $$PATH)
lint_confidence=0.9
min_coverage=72
path=./...
steps=0
op=up

-include config.env
export

setup: .make.setup
.make.setup:
	GO111MODULE=off go get -u github.com/kyoh86/richgo
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get -u golang.org/x/tools/cmd/cover
	GO111MODULE=off go get -u github.com/golang/mock/mockgen
	touch .make.setup


test: setup
	$(GOPATH)/bin/richgo test $(path)

lint: setup
	$(GOPATH)/bin/golint -set_exit_status -min_confidence $(lint_confidence) $(path)
	@echo "Golint found no problems on your code!"
