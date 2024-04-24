package main

import (
	pb "github.com/sajuno/pace-generic-decorators/api/proto"
	"github.com/sajuno/pace-generic-decorators/server"
	"github.com/sajuno/pace-generic-decorators/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.GRPCAuthInterceptor))
	pb.RegisterTodosServer(grpcServer, server.NewGRPCServer(store.NewMemoryStore()))

	log.Println("running todo grpc server on localhost:8080")
	log.Fatal(grpcServer.Serve(lis))
}
