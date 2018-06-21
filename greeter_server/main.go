package main

import (
	pb "github.com/mwaaas/grpc_tutorial/hello_world"
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"fmt"
)

// server is used to implement helloworld.GreeterServer.
type server struct {}


func (s *server) SayHello(ctx context.Context, request * pb.HelloRequest)  (*pb.HelloReply, error){
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}

func main(){
	lis, err := net.Listen("tcp", "0.0.0.0:80")

	if err != nil{
		log.Fatal("Failed to listen to %v ", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	fmt.Println("Server started:", "0.0.0.0:80")
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to server %v", err)
	}
}


