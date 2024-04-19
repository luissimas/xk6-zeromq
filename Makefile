GO = $(shell which go 2>/dev/null)

.PHONY: all clean xk6 build

all: clean format build

xk6:
	$(GO) install go.k6.io/xk6/cmd/xk6@latest

build: xk6
	xk6 build v0.50.0 --with github.com/luissimas/xk6-zeromq=. --output bin/k6

format:
	go fmt ./...

clean:
	$(RM) bin/k6
