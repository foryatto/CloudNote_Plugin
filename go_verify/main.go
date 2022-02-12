package main

import (
	"fmt"
	pb "go_verify/proto"
	. "go_verify/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = 50051

func main() {
	listen, err := net.Listen("tcp",fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterVerifyServer(s, &Server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
