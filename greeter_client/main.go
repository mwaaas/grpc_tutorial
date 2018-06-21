package main

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
	pb "github.com/mwaaas/grpc_tutorial/hello_world"
	"os"
	"time"
	"fmt"
)

func main()  {
	// setup connection
	conn, err := grpc.Dial("greet_server:80", grpc.WithInsecure())

	if err != nil{
		log.Fatal("unable to connect: ", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := "world"

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatal("Could not greet: ", err)
	}

	fmt.Println("Greeting: %s", r.Message)
}
