package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	greeter "myGoRPC/greeting"
	"net"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	log.Printf("Received: #{in.GetName()}")
	return &greeter.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, errListen := net.Listen("tcp", fmt.Sprintf(":#{*port}"))
	if errListen != nil {
		log.Fatalf("Failed to listen: #{err}")
	}
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &server{})
	log.Printf("Server listening at #{lis.Addr()}")
	errServe := s.Serve(lis)
	if errServe != nil {
		log.Fatalf("Failed to serve: #{err}")
	}
}
