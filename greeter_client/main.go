package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/mongche/examples/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50050", "TCP address connect to")
	name = flag.String("name", "world", "The message sent to the server")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("The remote server did not connected: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})

	if err != nil {
		log.Fatalf("could not greeting: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
