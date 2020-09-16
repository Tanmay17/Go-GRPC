package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Go-GRPC/calculator/calcpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calcpb.CalculateRequest) (*calcpb.CalculateResponse, error) {
	fmt.Printf("Calculate Function was invoked with %v", req)
	val1 := req.GetCalculator().GetFirstVal()
	val2 := req.GetCalculator().GetSecondVal()

	sum := val1 + val2

	res := &calcpb.CalculateResponse{
		Sum: sum,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
