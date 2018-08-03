package main

// grpc
// https://github.com/grpc/grpc-go

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "../snippets"
)

const dateFormat = "2006-01-02 15:04:05"

const imageUrl = "https://source.unsplash.com/1600x900?dog"

const (
	port = ":50051"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func printSayHellow(c pb.SnippetsClient, ctx context.Context, req *pb.SnippetRequest, cancel context.CancelFunc) {
	r, err := c.SayHello(ctx, req)
	if err != nil {
		log.Printf("could not greet: %v", err)
	} else {
		log.Printf("Greeting: %s", r.Message)
		cancel()
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSnippetsClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < 10; i++ {
		log.Printf("Print: %d", i)
		go printSayHellow(c, ctx, &pb.SnippetRequest{Name: fmt.Sprintf("%s %d", name, i), Sleep: int32(1000 * (i + 1))}, cancel)
	}
	time.Sleep(3 * time.Second)
	fmt.Print("Finished")
}
