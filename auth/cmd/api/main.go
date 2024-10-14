package main

import (
	"flag"
	"fmt"
	"net"

	"auth/internal/database"
	"auth/internal/server"

	pb "auth/proto"

	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 5001, "The server port")

func main() {
	flag.Parse()

	slog.Info("Connecting to Database")
	db, err := database.New()
	if err != nil {
		slog.Error(err)
	}

	slog.Info("Running Migrations")
	if db.Migrate(database.Accounts{}) != nil {
		slog.Error(err)
	}

	slog.Info("Starting Server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthenticationServer(s, &server.Server{})
	reflection.Register(s)

	slog.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		slog.Fatalf("failed to serve: %v", err)
	}
}
