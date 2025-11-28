package main

import (
	"context"
	"fmt"
	"log"
	"metadatasvc/gen/go/staticpagepb"
	"metadatasvc/internal/bootstrap"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := bootstrap.Init(); err != nil {
		panic(err)
	}

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", bootstrap.Config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	staticPageClient := staticpagepb.NewStaticPageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := staticPageClient.Create(ctx, &staticpagepb.CreateRequest{})
	fmt.Println("errrrrrrrr", err)
	fmt.Println("resp", resp)
}
