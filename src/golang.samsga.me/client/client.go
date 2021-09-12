package main

import (
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	pb "golang.samsga.me/apis/skeleton"
)

var (
	host = flag.String("host", "localhost", "host of the gRPC server")
	port = flag.Int("port", 4000, "port of the gRPC server")
)

func main() {
	flag.Parse()
	conn,err := grpc.Dial(fmt.Sprintf("%s:%d",*host,*port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(fmt.Sprintf("Error connecting to server: %v", err))
	}
	defer conn.Close()

	c := pb.NewSkeletonClient(conn)

	ctx := context.Background()
	r, err := c.Ping(ctx, &pb.PingRequest{Message: "Hello there"})
	if err != nil {
		panic("Error pinging server")
	}
	fmt.Printf("Response: %v\n", r)
}
