GO ?= go
BINARY = sudoku-solver

all: clean build test run

clean:
	$(GO) clean
	rm -rf bin

build:
	if [ ! -d bin ]; then mkdir bin; fi
	$(GO) build -v -o bin/$(BINARY)

test:
	$(GO) test -v ./sudoku

run:
	bin/$(BINARY) -i fixtures/sample_1.txt 