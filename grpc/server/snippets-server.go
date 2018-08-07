package main

// grpc
// https://github.com/grpc/grpc-go

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "../snippets"
)

const dateFormat = "2006-01-02 15:04:05"

const imageUrl = "https://source.unsplash.com/1600x900?dog"

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.SnippetRequest) (*pb.SnippetResponse, error) {
	logRequest(in)
	time.Sleep(time.Duration(in.Sleep) * time.Millisecond)
	return &pb.SnippetResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	fmt.Println("Run client")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSnippetsServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func logRequest(req *pb.SnippetRequest) {
	log.Println("tcp snippets request", req.Sleep, " "+time.Now().Format(dateFormat)) // TODO: Middleware
}
