package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewOutpatientHistoryServiceClient(conn)
	getOutpatientHistories(c)
}

func getOutpatientHistories(c pb.OutpatientHistoryServiceClient) {
	log.Println("---OutpatientHistories was invoked---")
	stream, err := c.GetOutpatientHistories(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling OutpatientHistories: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
