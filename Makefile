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

gz_releases=$(addsuffix .tar.gz, $(PLATFORM_LIST))

$(gz_releases): %.tar.gz : %
	tar czf $(NAME)-$(VERSION)-$@ -C $(NAME)-$(VERSION)-$</ ../LICENSE $(NAME)

sha256_releases=$(addsuffix .tar.gz.sha256, $(PLATFORM_LIST))

$(sha256_releases): %.sha256 : %
	shasum -a 256 $(NAME)-$(VERSION)-$< > $(NAME)-$(VERSION)-$@

releases: $(gz_releases) $(sha256_releases)

install:
	CGO_ENABLED=0 go install -trimpath -ldflags $(GO_LDFLAGS)