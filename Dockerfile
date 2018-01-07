FROM golang:1.9

# required to force cgo (for sqlite driver) with cross compile
ENV CGO_ENABLED 1

# i386 cross compilation
RUN dpkg --add-architecture i386 && \
	apt-get update && \
	apt-get install -y libc6-dev-i386 && \
	rm -rf /var/lib/apt/lists/*

# development dependencies
RUN go get \
	github.com/golang/lint/golint \
	github.com/kisielk/errcheck

# copy source files
COPY . $GOPATH/src/github.com/amacneil/dbmate
WORKDIR $GOPATH/src/github.com/amacneil/dbmate

# build
RUN go install -v ./cmd/dbmate

CMD dbmate
