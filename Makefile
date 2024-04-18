GO = $(shell which go 2>/dev/null)

.PHONY: all clean xk6 build

all: clean build

xk6:
	$(GO) install go.k6.io/xk6/cmd/xk6@latest

build: xk6
	xk6 build v0.50.0 --with github.com/luissimas/xk6-diameter=. --output bin/k6

clean:
	$(RM) bin/k6 bin/dict_generator
