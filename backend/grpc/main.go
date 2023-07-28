package main

import (
	context "context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	UnimplementedHelloWorldServiceServer
}

func (s *server) Say(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	return &HelloWorldResponse{
		Msg: req.Me + " say : hello world",
	}, nil
}

func main() {
	errc := make(chan error)

	lis, err := net.Listen("tcp", ":5033")
	if err != nil {
		errc <- err
	}

	s := grpc.NewServer()
	RegisterHelloWorldServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		errc <- s.Serve(lis)
	}()

	log.Fatal(<-errc)
}
