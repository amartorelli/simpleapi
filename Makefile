# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=simpleapi
CMD_PATH=./$(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
run-sudo:
	sudo ./$(BINARY_NAME)
run-sudo-debug:
	sudo ./$(BINARY_NAME) -loglevel="debug"
all: test build
build-run: build run-sudo
build-run-debug: build run-sudo-debug