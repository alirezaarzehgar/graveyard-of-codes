VERSION=$$(git rev-list -1 HEAD)

all: build

build:
	go build -ldflags "-X main.VERSION=${VERSION}"
