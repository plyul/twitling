FROM golang:1.15

ARG PROTOC_VERSION="3.14.0"
ARG GOLANGCI_LINT_VERSION="1.35.2"

ENV GOBIN=$GOPATH/bin

RUN go get github.com/bufbuild/buf/cmd/buf \
    github.com/bufbuild/buf/cmd/protoc-gen-buf-check-breaking \
    github.com/bufbuild/buf/cmd/protoc-gen-buf-check-lint \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/twitchtv/twirp/protoc-gen-twirp
RUN apt-get update && apt-get -y install unzip
RUN curl -sLO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /tmp/protoc && \
    cp /tmp/protoc/bin/protoc "$GOBIN/protoc" && rm -rf /tmp/protoc
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_LINT_VERSION}

RUN ls -la "$GOPATH"/bin
WORKDIR /app
