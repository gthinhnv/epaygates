package main

import (
	"apigateway/gen/go/staticpagepb"
	"apigateway/internal/bootstrap"
	"apigateway/internal/grpc/staticpageservice"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	httpServer := &http.Server{
		Handler: router,
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
