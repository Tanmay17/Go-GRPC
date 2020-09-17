package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Go-GRPC/greet/greetpb"

	"google.golang.org/grpc"
)

func doUnary(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Tanmay",
			LastName:  "Srivastava",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Tanmay",
			LastName:  "Srivastava",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading file: %v", err)
		}
		log.Printf("Response from Greet: %v", msg.GetResult())
	}
}

func main() {

	fmt.Println("Hello I'm Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
	doServerStreaming(c)
}
