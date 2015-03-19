package main

import (
	"log"
	"net"

	"github.com/calavera/git-explorer/backend/gitops"
	"github.com/calavera/git-explorer/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRepositoryExplorerServer(grpcServer, gitops.NewServer())

	log.Println("Backend server started in :9090")
	grpcServer.Serve(lis)
}
