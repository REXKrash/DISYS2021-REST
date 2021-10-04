package main

import (
	"context"
	"log"
	"net"

	pb "github.com/REXKrash/DISYS2021-REST/gRPC/routeguide"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCourseServiceServer
}

func (s *server) CreateCourse(ctx context.Context, in *pb.Course) (*pb.CourseResponse, error) {
	log.Printf("Received CreateCourse request with course name: %v", in.GetName())
	log.Println("Sending response...")
	return &pb.CourseResponse{Success: true, Message: "Successfully created course with name " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
