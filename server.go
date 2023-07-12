package main

import (
	"fmt"
	"net"

	"example.com/example-app/stock_service"
	"google.golang.org/grpc"
)

var (
	PORT = 5001
)

func main() {
	fmt.Printf("Starting the server at port %v\n", PORT)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	s := stock_service.Server{}
	grpcServer := grpc.NewServer()
	stock_service.RegisterStockServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Failed to server the gRPC server")
		return
	}

	fmt.Println("Successfully started the server!")
}
