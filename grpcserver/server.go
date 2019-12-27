package grpcserver

import (
	context "context"
	"fmt"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

type nameServer struct{}

func (ns *nameServer) GetName(ctx context.Context, wName *WithName) (*WithName, error) {
	return wName, fmt.Errorf("Not implemented")
}

func Serve() {
	lis, err := net.Listen("tcp", ":8088")
	fmt.Println("Serving grpc at 8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterGetNameServiceServer(grpcServer, &nameServer{})
	grpcServer.Serve(lis)
}
