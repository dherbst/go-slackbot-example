
all: src/chatty/chatty.go src/chatty/hellohandler.go src/chatty/messagehandler.go src/chatty/cmd/main.go
	echo $(PWD)
	GO15VENDOREXPERIMENT=1 GOPATH=$(PWD) go build -o chatty src/chatty/cmd/main.go


clean:
	rm chatty
