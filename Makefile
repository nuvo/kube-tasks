HAS_DEP := $(shell command -v dep;)
DEP_VERSION := v0.5.0

all: bootstrap build

fmt:
	go fmt ./pkg/... ./cmd/...

vet:
	go vet ./pkg/... ./cmd/...

# Build orca binary
build: fmt vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/kube-tasks cmd/kube-tasks.go

bootstrap:
ifndef HAS_DEP
	wget -q -O $(GOPATH)/bin/dep https://github.com/golang/dep/releases/download/$(DEP_VERSION)/dep-linux-amd64
	chmod +x $(GOPATH)/bin/dep
endif
	dep ensure
