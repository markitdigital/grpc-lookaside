LDFLAGS=-ldflags "-s -w"

.DEFAULT_GOAL := all
all: clean ensure proto test windows linux

clean:
	$(RM) -rf ./bin/*
	$(RM) -rf ./_proto/*.go

ensure:
	dep ensure

linux:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ./bin/windows_amd64/grpc-lookaside.exe main.go

proto:
	protoc --go_out=plugins=grpc:. ./_proto/lookaside.proto

test:
	go test -cover

windows:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./bin/linux_amd64/grpc-lookaside main.go