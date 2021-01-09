package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tpc", ":8080")
	if err != nil {
		log.Fatalf("FAILED %v", err)
	}
	grcpServer := grpc.NewServer()

	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed : %s", err)
	}
}
