package main

import (
	"context"
	"log"
	"time"

	pb "github.com/REXKrash/DISYS2021-REST/gRPC/routeguide"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer connection.Close()
	client := pb.NewCourseServiceClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Println("Sending request to create course...")
	response, err := client.CreateCourse(ctx, &pb.Course{Id: 5, Name: "DISYS"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Response: %s", response.GetMessage())
}
