package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/dusk-chancellor/go-grpc-examples/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GeometryServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Area(
	ctx context.Context,
	in *pb.RectRequest,
) (*pb.AreaResponse, error) {
	log.Println("Invoked Area:", in)

	return &pb.AreaResponse{
		Result: in.Height * in.Width,
	}, nil
}

func (s *Server) Perimeter(
	ctx context.Context,
	in *pb.RectRequest,
) (*pb.PerimeterResponse, error) {
	log.Println("Invoked Perimeter: ", in)

	return &pb.PerimeterResponse{
		Result: 2 * (in.Height + in.Width),
	}, nil
}

func main() {
	host := "localhost"
	port := "5000"

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
	log.Println("tcp listener started at", port)

	grpcServer := grpc.NewServer()
	geomServiceServer := NewServer()
	pb.RegisterGeometryServiceServer(grpcServer, geomServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error %v", err)
		os.Exit(1)
	}
}
