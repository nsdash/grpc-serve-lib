# Coverage for "google.golang.org/grpc"

___

## Usage

* Create .env file and set params
```
NETWORK=tcp
NETWORK_PORT=9000
```

* Use Manager
```go
  protoMessageHandler := message_handler_factory.NewProtoMessageHandlerFactory().MakeMessageHandler()

  grpcServer := server.NewGrpcServer()

  grpcServer.RunServer(func(server *grpc.Server) {
    proto.RegisterProtoMessageHandlerServer(server, protoMessageHandler)
  })
  ```
