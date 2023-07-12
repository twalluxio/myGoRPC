package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	greeter "myGoRPC/greeting"
	"net"
)

type server struct {
}

const port = ":50051"

func (s *server) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	log.Println("Received:", in.GetName())
	return &greeter.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Failed to listen at localhost" + port)
		return
	}
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &server{})
	log.Println("Server listening at localhost" + port)
	if err := s.Serve(lis); err != nil {
		log.Println("Failed to serve ", err)
		return
	}
}
