LDFLAGS=-ldflags "-s -w"
.DEFAULT_GOAL := all
all: clean ensure proto test windows linux zip

clean:
	$(RM) -rf ./bin/*
	$(RM) -rf ./_proto/*.go

docker:
	docker build -t grpc-lookaside:latest .

ensure:
	dep ensure

linux:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ./bin/windows_amd64/grpc-lookaside.exe ./cmd/grpc-lookaside.go

proto:
	protoc --go_out=plugins=grpc:. ./_proto/lookaside.proto

test:
	go test -cover github.com/markitondemand/grpc-lookaside

windows:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./bin/linux_amd64/grpc-lookaside ./cmd/grpc-lookaside.go

zip:
	7z a bin/grpc-lookaside-linux-amd64.zip bin/linux_amd64/
	7z a bin/grpc-lookaside-windows-amd64.zip bin/windows_amd64/