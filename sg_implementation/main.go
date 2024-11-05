package main

import (
	"context"
	"fmt"

	greeter "github.com/dusk-chancellor/go-grpc-examples/sg_implementation/api/gen/v1"
	rkboot "github.com/rookie-ninja/rk-boot"
	rkgrpc "github.com/rookie-ninja/rk-grpc/boot"
	"google.golang.org/grpc"
)

func main() {
	// creating instance `rkboot.Boot()``
	boot := rkboot.NewBoot()

	// adding entry `greeter`
	grpcEntry := boot.GetEntry("greeter").(*rkgrpc.GrpcEntry)

	// registering GRPC
	grpcEntry.AddRegFuncGrpc(registerGreeter)

	// adding endpoint
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// booting
	boot.Bootstrap(context.Background())

	boot.WaitForShutdownSig(context.Background())
}

// Registering GRPC
func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

// GRPC struct implementation
type GreeterServer struct {
	greeter.UnimplementedGreeterServer
}

// GRPC Handler Function
func (server *GreeterServer) Greeter(ctx context.Context, request *greeter.GreeterRequest) (*greeter.GreeterResponse, error) {
    return &greeter.GreeterResponse{
        Message: fmt.Sprintf("Hello %s!", request.Name),
    }, nil
}
