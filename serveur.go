package main

import (
	"context"
	"log"
	"net"

	pb "github.com/eya20/LogName/personpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *server) SendPerson(ctx context.Context, req *pb.PersonRequest) (*pb.PersonResponse, error) {
	log.Printf("Received: %s %s", req.Name, req.Surname)
	return &pb.PersonResponse{Message: "Received: " + req.Name + " " + req.Surname}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
