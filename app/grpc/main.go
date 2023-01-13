package main

import (
	"log"
	"net"

	pb "github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.OutpatientHistoryServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterOutpatientHistoryServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
