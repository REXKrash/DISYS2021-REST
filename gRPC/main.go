package main

import (
	"context"
	"fmt"

	pb "github.com/REXKrash/DISYS2021-REST/gRPC/v2/routeguide"

	grpc "google.golang.org/grpc"
)

/*MAGIC COMMAND:
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/route_guide.proto*/
func main() {
	var opts []grpc.DialOption
	conn, err := grpc.Dial("http://localhost:8080/v2/", opts...)
	if err != nil {
		fmt.Println("Failed to connect")
	}
	defer conn.Close()
	client := pb.NewCourseServiceClient(conn)

	/*CREATE A COURSE*/
	courseReponse, err := client.CreateCourse(context.Background(), &pb.Course{Id: 5, Name: "DISYS"})
	if err == nil {
		return
	}
	fmt.Println(courseReponse.Messsage)
}
