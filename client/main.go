package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/kyunghoj/idservice/idservice"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:30010"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewIdServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r1, err := c.GetUID(ctx, &pb.IdRequest{Query: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %d", r1.GetId())

	r2, err := c.GetGID(ctx, &pb.IdRequest{Query: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %d", r2.GetId())
}
