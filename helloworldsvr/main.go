package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":500051"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("received: %v", in.Name)
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
