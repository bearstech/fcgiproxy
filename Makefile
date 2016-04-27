.PHONY: linux build clean linux-src

linux:
	docker run -it --rm -v `pwd`:/go golang:1.6 go build -v -o fcgiproxy

linux-src:
	docker run -it --rm -v `pwd`:/go golang:1.6 go get

src:
	go get

build: src
	go build

clean:
	rm -f fcgiproxy
	rm -rf src
