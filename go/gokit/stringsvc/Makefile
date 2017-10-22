export GOPATH = $(PWD)

all: build

dep:
	cd src/stringsvc/ && dep ensure

build: dep
	go install stringsvc/

test:
