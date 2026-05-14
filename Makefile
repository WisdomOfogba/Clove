SERVER_DIR = "./"

.PHONY: build-server
build:
	@go build -o bin/server $(SERVER_DIR)
prod:
	@go build -ldflags "-w -s" -o bin/server $(SERVER_DIR)
