package main

import (
	"context"
	"errors"
	"fmt"
	proto "grpcc/protoc"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExmapleServer
}

func main() {
	fmt.Println("Welcome to the server")
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		fmt.Println("Failed to start the server")
		panic(tcpErr)
	}

	srv := grpc.NewServer() //engine
	proto.RegisterExmapleServer(srv, &server{})
	reflection.Register(srv)

	err := srv.Serve(listener)
	if err != nil {
		fmt.Println("Failed to start the server 2")
		log.Fatal(err)
		return
	}
	fmt.Println("Grpc server successfully started on port :9000")

}

func (s *server) ServerReply(c context.Context, req *proto.HellpRequest) (*proto.HelloResponse, error) {
	fmt.Println("receive req from client:", req.SomeString)
	res:= &proto.HelloResponse{Reply: "Received Successfully"}
	return res, errors.New("")
}

// Method to start grpc server:
// build a listener
//form a new server
// register that server
