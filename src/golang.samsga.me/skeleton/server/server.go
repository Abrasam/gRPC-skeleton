package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "golang.samsga.me/apis/skeleton"
)

var (
	port = flag.Int("port", 4000, "the port to serve gRPC requests to")
)

type server struct {
	pb.UnimplementedSkeletonServer
}

func main() {
	flag.Parse()
	sock,err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d",*port))
	if err != nil {
		panic("Failed to start server")
	}
	s := grpc.NewServer()
	pb.RegisterSkeletonServer(s, &server{})
	fmt.Printf("Listening on %v\n", sock.Addr())
	if err := s.Serve(sock); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

/*
 * gRPC Functions
 */

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.Pong, error) {
	fmt.Printf("Ping!\n")
	return &pb.Pong{Message: req.Message},nil
}
