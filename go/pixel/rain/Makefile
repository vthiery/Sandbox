export GOPATH = $(PWD)
GO_BUILD_OPTS =

all: build

dep:
	cd src/rain/ && dep ensure

build: dep
	go install $(GO_BUILD_OPTS) rain/