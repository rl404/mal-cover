# Golang base image
FROM golang:1.15 as go_builder
WORKDIR /go/src/github.com/rl404/mal-cover
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o mal-cover

# New stage from scratch
FROM alpine:3.13
RUN apk add --no-cache ca-certificates
WORKDIR /docker/bin
COPY --from=go_builder /go/src/github.com/rl404/mal-cover/mal-cover mal-cover
CMD ["/docker/bin/mal-cover"]