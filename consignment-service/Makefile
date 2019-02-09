
build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/xmarlem/shippy/consignment-service proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run --rm -d -p 50051:50051 --name cons-srv consignment-service

stop:
	docker stop cons-srv