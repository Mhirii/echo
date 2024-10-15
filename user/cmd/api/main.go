package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"user/internal/database"
	"user/internal/server"
	pb "user/proto"
)

var (
	port         = flag.Int("port", 5002, "The server port")
	dependencies = []string{"AUTH_ADDR"}
)

func main() {
	flag.Parse()

	slog.Info("Checking Services")
	addrs, err := getServiceAddresses(dependencies)
	if err != nil {
		slog.Error(err)
		panic(err)
	}
	err = checkServices(addrs)
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	slog.Info("Connecting to Database")
	db, err := database.New()
	if err != nil {
		slog.Error(err)
	}

	slog.Info("Running Migrations")
	if db.Migrate(database.User{}) != nil {
		slog.Error(err)
	}

	slog.Info("Starting Server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server.Server{AuthAddr: addrs["AUTH_ADDR"]})
	reflection.Register(s)

	slog.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		slog.Fatalf("failed to serve: %v", err)
	}
}

func getServiceAddr(addrKey string) (string, error) {
	addr := os.Getenv(addrKey)
	if addr == "" {
		return "", errors.New(addrKey + " is not set")
	}
	return addr, nil
}

func getServiceAddresses(dependencies []string) (map[string]string, error) {
	services := make(map[string]string)
	for _, dependency := range dependencies {
		addr, err := getServiceAddr(dependency)
		if err != nil {
			return nil, err
		}
		services[dependency] = addr
	}
	return services, nil
}

func checkServices(services map[string]string) error {
	for service, addr := range services {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return errors.New(service + " is not reachable at " + addr)
		}
		conn.Close()
	}
	return nil
}
