linux:
	#docker run -it --rm -v `pwd`:/go golang:1.6 go get
	docker run -it --rm -v `pwd`:/go golang:1.6 go build -v -o fcgiproxy
