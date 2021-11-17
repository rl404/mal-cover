# Golang base image
FROM golang:1.17 as go_builder
LABEL stage="mal_cover_builder"
WORKDIR /mal-cover
COPY . .
WORKDIR /mal-cover/cmd/mal-cover
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o mal-cover

# New stage from scratch
FROM alpine:3.14
LABEL stage="mal_cover_binary"
RUN apk add --no-cache ca-certificates
COPY --from=go_builder /mal-cover/cmd/mal-cover/mal-cover /cmd/mal-cover/mal-cover
CMD ["/cmd/mal-cover/mal-cover", "server"]