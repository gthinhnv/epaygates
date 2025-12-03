package main

import (
	"cms/gen/go/staticpagepb"
	"cms/internal/bootstrap"
	"cms/internal/grpc/staticpageservice"
	"cms/internal/http/router"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
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

	// cmux for protocol multiplexing
	m := cmux.New(lis)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	// --- gRPC server ---
	grpcServer := grpc.NewServer()
	staticpagepb.RegisterStaticPageServiceServer(grpcServer, staticpageservice.NewStaticPageServiceServer())
	go func() {
		log.Printf("gRPC server listening on port %d", bootstrap.Config.Port)
		if err := grpcServer.Serve(grpcL); err != nil {
			panic(err)
		}
	}()

	// --- Gin HTTP server ---
	r := router.New()

	httpServer := &http.Server{
		Handler: r,
	}

	go func() {
		log.Printf("HTTP server (Gin) listening on port %d", bootstrap.Config.Port)
		if err := httpServer.Serve(httpL); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Start cmux
	log.Println("Starting cmux multiplexer...")
	if err := m.Serve(); err != nil {
		panic(err)
	}
}
