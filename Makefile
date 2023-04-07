GO_PACKAGES ?= $(shell go list ./...)

run:
	go run main.go

test:
	@go test -race -v ${GO_PACKAGES}
