package main

import (
	"fmt"
	"log"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/mapper"
	pb "github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/proto"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*Server) GetOutpatientHistories(_ *emptypb.Empty, stream pb.OutpatientHistoryService_GetOutpatientHistoriesServer) error {
	log.Println("OutpatientHistories was invoked")

	service := service.HospitalService{Client: &client.CoronaClient{}}
	outpatinet_histories, err := service.GetOutpatientHistories()
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	for _, outpatinet_history := range outpatinet_histories {
		err := stream.Send(mapper.OutpatientHistoryEntityToPbOutpatientHistory(outpatinet_history))

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Unknown internal error: %v", err),
			)
		}

	}

	return nil
}
