GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
# BINARY_NAME=go-fetch
BINARY_PATH=~/bin/

default:
	$(GOBUILD) -o $(BINARY_PATH)
	