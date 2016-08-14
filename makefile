
all: src/chatty/chatty.go src/chatty/hellohandler.go src/chatty/messagehandler.go src/chatty/cmd/main.go
	echo $(PWD)
	GO15VENDOREXPERIMENT=1 GOPATH=$(PWD) go build -o chatty src/chatty/cmd/main.go


# go get does not place things into the vendor directory, we have to move it there manually
get:
	mkdir -p src/vendor
	GO15VENDOREXPERIMENT=1 GOPATH=$(PWD) go	get -u golang.org/x/net/websocket
	mv src/golang.org src/vendor/golang.org


clean:
	rm chatty
