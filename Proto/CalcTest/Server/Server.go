package main

import (
	"context"
	"fmt"
	pb "github.com/Jeremytjuh/Gotraining/Proto/CalcTest/CalcTest"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9090"
)

type server struct{}

func (s *server) Calculate(ctx context.Context, in *pb.CalcRequest) (*pb.CalcReply, error) {
	message, err := Calculate(in)
	return &pb.CalcReply{Message: fmt.Sprint(message)}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculateServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
