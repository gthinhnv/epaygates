package staticpageservice

import (
	"apigateway/gen/go/staticpagepb"
	"apigateway/internal/bootstrap"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StaticPageServiceServer struct {
	staticpagepb.UnimplementedStaticPageServiceServer
}

func NewStaticPageServiceServer() *StaticPageServiceServer {
	return &StaticPageServiceServer{}
}

func (s *StaticPageServiceServer) Create(ctx context.Context, req *staticpagepb.CreateRequest) (*staticpagepb.CreateResponse, error) {
	conn, err := grpc.NewClient(bootstrap.SharedConfig.MetadataService.GRPC.LocalAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	staticPageClient := staticpagepb.NewStaticPageServiceClient(conn)

	resp, err := staticPageClient.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &staticpagepb.CreateResponse{
		Id: resp.Id,
	}, nil
}
