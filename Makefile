.PHONY: build run test lint clean

build:
	go build -o github-activity

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -f github-activity