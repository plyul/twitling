FROM golang:1.15

ARG PROTOC_VERSION="3.14.0"

ENV GOBIN="$GOPATH/bin"

RUN go get github.com/bufbuild/buf/cmd/buf \
    github.com/bufbuild/buf/cmd/protoc-gen-buf-check-breaking \
    github.com/bufbuild/buf/cmd/protoc-gen-buf-check-lint \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/twitchtv/twirp/protoc-gen-twirp
RUN apt update && apt install -y unzip
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /tmp/protoc && \
    cp /tmp/protoc/bin/protoc "$GOBIN/protoc"
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0

RUN ls -la "$GOPATH"/bin
WORKDIR /app
ADD build.sh /
