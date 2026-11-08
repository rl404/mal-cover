# Golang base image
FROM golang:1.27.3 AS go_builder
LABEL stage="mal_cover_builder"
WORKDIR /mal-cover
COPY . .
WORKDIR /mal-cover/cmd/mal-cover
RUN CGO_ENABLED=0 go build -mod=vendor -ldflags="-s -w" -o mal-cover

# New stage from scratch
FROM alpine:3.24.3
LABEL stage="mal_cover_binary"
RUN apk add --no-cache ca-certificates && adduser -D -u 10001 mal-cover
COPY --from=go_builder /mal-cover/cmd/mal-cover/mal-cover /usr/local/bin/mal-cover
WORKDIR /home/mal-cover
USER mal-cover
ENTRYPOINT ["/usr/local/bin/mal-cover"]
