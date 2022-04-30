## gRPC Exmpales
Quick start gRPC examples with a simple working example.

Make sure you've installed go programming language and protocol plugins:
```console
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Add a protobuf file (called `main.proto`) and generate to the binary code via the following command.
```console
$ ./build.sh
```

Run the client and server with the following command
1. Set up the server:
```console
$ go run greeter_server/main.go
```

2. From another window, run a command with the arguemnt:
```console
$ go run greeter_client/main.go --name=Bob

// Print the follwoing output
Greeting: Hello Bob
Greeting: Hello again Bob
```
