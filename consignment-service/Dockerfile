# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from golang v1.11 base image
FROM golang:1.11 as builder

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/xmarlem/shippy/consignment-service

# Copy everything from the current directory to the PWD inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/consignment-service .


######## Start a new stage from scratch ########

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /go/bin/consignment-service .

EXPOSE 50051

ENTRYPOINT ["./consignment-service"]