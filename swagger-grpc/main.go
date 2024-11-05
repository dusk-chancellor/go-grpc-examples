package main

import (
	"context"

	"github.com/dusk-chancellor/go-grpc-examples/swagger-grpc/gen/go/proto/product"
	"github.com/dusk-chancellor/go-grpc-examples/swagger-grpc/gen/go/proto/user"
	rkboot "github.com/rookie-ninja/rk-boot"
	rkgrpc "github.com/rookie-ninja/rk-grpc/boot"
	"google.golang.org/grpc"
)

// http://localhost:50051/sw

func main() {
	boot := rkboot.NewBoot()

    grpcEntry := boot.GetEntry("swagger-grpc").(*rkgrpc.GrpcEntry)
    
    grpcEntry.AddRegFuncGrpc(registerProductService, registerUserService)
	grpcEntry.AddRegFuncGw(product.RegisterProductServiceHandlerFromEndpoint, user.RegisterUserServiceHandlerFromEndpoint)

    boot.Bootstrap(context.Background())

    boot.WaitForShutdownSig(context.Background())
}

type ProductServer struct {
	product.UnimplementedProductServiceServer
}

type UserServer struct {
	user.UnimplementedUserServiceServer
}

func registerProductService(server *grpc.Server) {
    product.RegisterProductServiceServer(server, ProductServer{})
}

func registerUserService(server *grpc.Server) {
    user.RegisterUserServiceServer(server, UserServer{})
}

