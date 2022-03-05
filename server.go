package main

import (
	"calculator-server/api"
	"calculator-server/calc"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	GRPC_HOST := os.Getenv("GRPC_HOST")
	GRPC_PORT := os.Getenv("GRPC_PORT")
	serverInfo := fmt.Sprintf("%s:%s", GRPC_HOST, GRPC_PORT)
	listen, err := net.Listen("tcp", serverInfo)
	if err != nil {
		log.Fatal("Failed to listen : ", err)
	}

	s := calc.Server{}

	grpcServer := grpc.NewServer()
	api.RegisterCalculatorServer(grpcServer, &s)

	// GRPC_HOST=localhost GRPC_PORT=9000 go run server.go
	log.Println("server runs on", serverInfo)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("failed to server", err)
	}
}
