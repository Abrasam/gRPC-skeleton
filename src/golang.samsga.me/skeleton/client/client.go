package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	pb "golang.samsga.me/apis/skeleton"
	"golang.samsga.me/shortcuts/utils"
)

var (
	host = flag.String("host", "localhost", "host of the gRPC server")
	port = flag.Int("port", 4000, "port of the gRPC server")
)

func main() {
	flag.Parse()
	conn,err := grpc.Dial(fmt.Sprintf("%s:%d",*host,*port), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))
	utils.PanicOnError("Failed to connect", err)
	defer conn.Close()

	c := pb.NewSkeletonClient(conn)

	ctx := context.Background()
	r, err := c.Ping(ctx, &pb.PingRequest{Message: "Hello there"})
	utils.PanicOnError("Failed to ping server", err)
	fmt.Printf("Response: %v\n", r)
}
