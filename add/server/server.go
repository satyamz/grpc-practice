package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/satyamz/grpc-practice/add"
)

const (
	port = ":58888"
)

//Server type for implementation of server
type Server struct{}

//Add returns addition of two numbers
func (s *Server) Add(ctx context.Context, numbers *pb.Numbers) (*pb.Addition, error) {
	number1 := numbers.Number1
	number2 := numbers.Number2
	addition := number1 + number2
	return &pb.Addition{
		Addition: addition,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdderServer(s, &Server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
