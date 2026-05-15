APP_NAME=dive
CMD_PATH=./cmd/dive

.PHONY: all build run install

all: build

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)

install:
	go install $(CMD_PATH)
