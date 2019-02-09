
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/xmarlem/shippy/consignment-service proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run --rm -d -p 50051:50051 --name cons-srv \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns consignment-service

stop:
	docker stop cons-srv