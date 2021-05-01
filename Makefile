FILE := sample.csv
HEADER := true

.PHONY: run
run:
	go run cmd/csview/main.go -f=$(FILE) -h=$(HEADER)

.PHONY: build
build:
	go build -o csview cmd/csview/main.go

.PHONY: test
test:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: mod
mod:
	go mod download