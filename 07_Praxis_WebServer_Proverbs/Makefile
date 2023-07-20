.PHONY: build fmt lint vet
.DEFAULT_GOAL: build

BINARY = "proserve"

fmt:
	go fmt .

build: clean
	go build -o $(BINARY) .

run:
	go run .

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - clean, fmt, build"
	@echo "make build - just build"
	@echo "make run - just run"
