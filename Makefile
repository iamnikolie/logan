.PHONY: all link build clean

export GOPATH=${PWD}/vendor:/usr/share/gocode:$(shell pwd)
export HOME=${PWD}
export GO111MODULE=off

all: build

link: clean
	go get -d -v github.com/go-chi/chi 

build: link
	go build -o bin/logan-server ./backend/src/main.go 

clean:
	rm -rf bin/logan
