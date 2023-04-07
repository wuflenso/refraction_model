GO_PACKAGES ?= $(shell go list ./...)

run:
	go run ./examples/trace_ray_refraction.go

test:
	@go test -race -v ${GO_PACKAGES} -coverprofile cover.out 
