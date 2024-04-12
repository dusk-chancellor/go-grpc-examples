package main

import (
	"context"
	"fmt"
	"os"

	pb "github.com/dusk-chancellor/go-gprc-examples/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewExpressionServiceClient(conn)

	ctx := context.TODO()
	ans1, err := grpcClient.Calculate(ctx, &pb.ExpressionRequest{
		Expression: "2+2*2",
	})
	if err != nil {
		fmt.Println("failed: ", err)
	}

	fmt.Printf("2+2*2 = %v", ans1.Result)
}