NAME ?= kubectl-nodepool
VERSION ?= $(shell git describe --tags || echo "unknown")
GO_LDFLAGS='-w -s'
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags $(GO_LDFLAGS)

PLATFORM_LIST = \
	darwin-amd64 \
	darwin-arm64 \
	linux-amd64 \
	linux-arm64

all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64

darwin-%:
	GOARCH=$* GOOS=darwin $(GOBUILD) -o $(NAME)-$(VERSION)-$@/$(NAME)

linux-%:
	GOARCH=$* GOOS=linux $(GOBUILD) -o $(NAME)-$(VERSION)-$@/$(NAME)

install:
	CGO_ENABLED=0 go install -trimpath -ldflags $(GO_LDFLAGS)