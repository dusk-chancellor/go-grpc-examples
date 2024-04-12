package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/dusk-chancellor/go-gprc-examples/calculator/internal/calculation"
	"github.com/dusk-chancellor/go-gprc-examples/calculator/internal/infix_to_postfix"
	pb "github.com/dusk-chancellor/go-gprc-examples/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.ExpressionServiceServer
}

func NewServer() *Server {
	return &Server{}
}

type ExpressionServiceServer interface {
	Calculate(ctx context.Context, in *pb.ExpressionRequest) (*pb.ExpressionResponse, error)
	mustEmbedUnimplementedGeometryServiceServer()
}

func (s *Server) Calculate(ctx context.Context, in *pb.ExpressionRequest) (*pb.ExpressionResponse, error) {
	postfixed := infix_to_postfix.ToPostfix(in.Expression)
	res, err := calculation.Evaluate(postfixed)
	if err != nil {
		return nil, err
	}
	return &pb.ExpressionResponse{Result: float32(res)}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	fmt.Println("tcp listener started")

	grpcServer := grpc.NewServer()

	exprServiceServer := NewServer()

	pb.RegisterExpressionServiceServer(grpcServer, exprServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}
