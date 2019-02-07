## gRPC

* High performance, open-source RPC framework
* Open source version of Google's internal framework Stubby
* Uses Protocol Buffers
* HTTP/2 for transport
* Bi-Directional streaming
* RPC is Efficient, Domain Specific and Strongly Typed
* Works across languages and platforms



### ProtoBuf
* Google's language-neutral, platform neutral, extensible mechanism for serialising structured data
* IDL - Describe once and generate interfaces for multiple languages
* Structure of the REquest and REsponse
* Binary format for network transmission
* Support multiple languages


### Types of RPC Methods
* Simple RPC
* Server-side streaming RPC
* Client-side streaming RPC
* Bi-directional streaming RPC

### Generate code for go grpc
```
protoc --go_out=plugins=grpc
```

### Go gRPC
* Install protoc compiler (Exceutable name is protoc)
* Install Go packages:
    * github.com/golang/protobuf/protoc
    * github.com/golang/protobuf/protoc-gen-go
    * google.golang.org/grpc

### grpc Interceptors
