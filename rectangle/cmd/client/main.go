package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/dusk-chancellor/go-grpc-examples/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	host := "localhost"
	port := "5000"
	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewGeometryServiceClient(conn)
	area, err := grpcClient.Area(context.TODO(), &pb.RectRequest{
		Height: 2,
		Width:  6,
	})
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	perim, err := grpcClient.Perimeter(context.TODO(), &pb.RectRequest{
		Height: 13,
		Width:  15,
	})
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Area:", area.Result)
	fmt.Println("Perimeter:", perim.Result)
}
