package main

import (
	"context"
	"fmt"
	proto "grpcc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Welcome to the client")
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed")
		panic(err)
	}

	client := proto.NewExmapleClient(conn)
	req := &proto.HellpRequest{SomeString: "hello from client"}
    client.ServerReply(context.TODO(),req)
}
