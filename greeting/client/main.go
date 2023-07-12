package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	greeter "myGoRPC/greeting"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var (
	name = flag.String("name", defaultName, "Input your name")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect failed on" + address)
	}
	defer conn.Close()
	c := greeter.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, err := c.SayHello(ctx, &greeter.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("Could not greet")
	}
	log.Printf("Greeting: " + r.GetMessage())
}
