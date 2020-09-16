package main

import (
	"context"
	"log"

	"github.com/Go-GRPC/calculator/calcpb"
	"google.golang.org/grpc"
)

func doUnary(c calcpb.CalculateServiceClient) {
	req := &calcpb.CalculateRequest{
		Calculator: &calcpb.Calculator{
			FirstVal:  17,
			SecondVal: 21,
		},
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while getting Response: %v", err)
	}
	log.Printf("Sum is: %v", res.Sum)
}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error while making connection")
	}

	defer conn.Close()
	c := calcpb.NewCalculateServiceClient(conn)

	doUnary(c)
}
