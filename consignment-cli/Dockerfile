FROM golang:1.11 as builder

WORKDIR /go/src/github.com/xmarlem/shippy/consignment-cli

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/consignment-cli .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY consignment.json /app/consignment.json

COPY --from=builder /go/bin/consignment-cli .


CMD ["./consignment-cli"]