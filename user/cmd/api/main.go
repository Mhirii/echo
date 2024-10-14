package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/gookit/slog"
	"google.golang.org/grpc"

	"user/internal/server"
	pb "user/proto"
)

var port = flag.Int("port", 50051, "The server port")

func main() {
	flag.Parse()
	slog.Info("Connecting to Database")

	slog.Info("Starting Server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisteruserServer(s, &server.Server{})

	slog.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		slog.Fatalf("failed to serve: %v", err)
	}
}

