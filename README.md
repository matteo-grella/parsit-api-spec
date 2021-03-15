# parsit-api-spec

This repository contains the OpenAPI and gRPC specifications for the Parsit dependency parsing service.

### API Development

**Requirements**:

* [Go 1.16](https://golang.org/dl/)
* [Go modules](https://blog.golang.org/using-go-modules)

Clone this repo or get the module:

```console
go get -u https://github.com/matteo-grella/parsit-api-spec
```

**Setup:**

```console
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```

```console
brew install protobuf
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/NYTimes/openapi2proto/cmd/openapi2proto
```

```console
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

**Generate API source:**

```console
go generate ./...
```