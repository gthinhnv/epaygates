package main

import (
	"fmt"
	"log"
	"metadatasvc/gen/go/staticpagepb"
	"metadatasvc/internal/bootstrap"
	"metadatasvc/internal/grpc/staticpageservice"
	"net"

	"google.golang.org/grpc"
)

func main() {
	if err := bootstrap.Init(); err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", bootstrap.Config.Port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()

	// Register services
	staticpagepb.RegisterStaticPageServiceServer(grpcServer, staticpageservice.NewStaticPageServiceServer())

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
