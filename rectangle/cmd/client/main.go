package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/dusk-chancellor/go-grpc-examples/rectangle/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	host := "localhost"
	port := "5000"

	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}

	defer conn.Close()

	grpcClient := pb.NewGeometryServiceClient(conn)

	area, err := grpcClient.Area(context.TODO(), &pb.RectRequest{
		Height: 10.1,
		Width:  20.5,
	})
	
	if err != nil {
		log.Println("failed invoking Area: ", err)
	}
	
	perim, err := grpcClient.Perimeter(context.TODO(), &pb.RectRequest{
		Height: 10.1,
		Width:  20.5,
	})
	
	if err != nil {
		log.Println("failed invoking Area: ", err)
	}
	
	fmt.Println("Area: ", area.Result)
	fmt.Println("Perimeter: ", perim.Result)
	
}