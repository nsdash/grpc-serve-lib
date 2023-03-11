package servers

import (
  "github.com/nsdash/config-lib"
  "google.golang.org/grpc"
  "net"
)

type GrpcServer struct {
  server        *grpc.Server
  configManager config.Manager
}

func NewGrpcServer() GrpcServer {
  return GrpcServer{
    configManager: config.NewManager(),
  }
}

func (ps *GrpcServer) RunServer(registerHandlerCallback func(server *grpc.Server)) {
  server := grpc.NewServer()

  ps.server = server

  registerHandlerCallback(server)

  listener := ps.makeListener()

  if err := server.Serve(listener); err != nil {
    panic(err)
  }
}

func (ps *GrpcServer) GracefulStop() {
  if ps.server == nil {
    panic("server is not defined")
  }

  ps.server.GracefulStop()
}

func (ps *GrpcServer) makeListener() net.Listener {
  listenerNetwork := ps.configManager.Get("NETWORK")
  listenerPort := ps.configManager.Get("NETWORK_PORT")

  listener, err := net.Listen(listenerNetwork, ":"+listenerPort)

  if err != nil {
    panic(err)
  }

  return listener
}
