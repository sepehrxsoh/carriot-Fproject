package main

import (
	"context"
	"encoding/json"
	"github.com/sepehrxsoh/carriot-Fproject/proto"
	"google.golang.org/grpc"
	"github.com/sepehrxsoh/carriot-Fproject/DataCr"
	"google.golang.org/grpc/reflection
	"net"

)
vehicles := DataCr.VehicleList

func (s *server) makelist(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	pay := request.GetPayload()
	x := int64(len(vehicles))
	if pay > x {
		pay = x
	}
	list := vehicles[:pay]
	result, _ := json.Marshal(list)
	return &proto.Response{Result: string(result)}, nil
}

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterListCustomersServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
