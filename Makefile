APP_NAME=dive
CMD_PATH=./cmd/dive

.PHONY: all build run install

all: build


GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION   ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

build:
	go build -ldflags "-X 'github.com/mzulfanw/dive/internal/cli.version=$(VERSION)' -X 'github.com/mzulfanw/dive/internal/cli.commit=$(GIT_COMMIT)' -X 'github.com/mzulfanw/dive/internal/cli.buildDate=$(BUILD_DATE)'" -o bin/$(APP_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)

install:
	go install $(CMD_PATH)
