.PHONY: default clean deps update-deps pretty test coverage coverage-profile coverage-html

default: clean deps pretty build test

clean:
	rm -rf _coverage

deps:
	go mod download

update-deps:
	go get -u ./...
	go mod tidy

pretty:
	gofmt -l -s -w .

build:
	go build -v ./...

test:
	go test ./...

coverage:
	go test -cover ./...

_coverage:
	mkdir _coverage

coverage-profile: _coverage
	go test -coverprofile=_coverage/coverage.out ./...

coverage-html: coverage-profile
	go tool cover -html=_coverage/coverage.out -o _coverage/coverage.html
