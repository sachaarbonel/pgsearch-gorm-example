# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)

# Stage 1 (to create a "build" image, ~850MB)
FROM golang:1.8-alpine AS builder
RUN go version

# Package dependencies
RUN apk add --update --no-cache git libc-dev

# install dependency tool
RUN go get github.com/golang/dep && go install github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/Sach97/pgsearch-gorm-example
COPY . /go/src/github.com/Sach97/pgsearch-gorm-example

# Get dependencies
RUN dep ensure -v

ENTRYPOINT ["go"]
CMD ["run","main.go"]
