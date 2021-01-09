package main

import (
	"google.golang.org/grcp"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tpc", ":8080")
	if err != nil {
		log.Fatalf("FAILED %v", err)
	}
	grcpServer := grcp.NewServer()

	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed : %s", err)
	}
}
