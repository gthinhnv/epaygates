package staticpageservice

import (
	"apigateway/gen/go/staticpagepb"
	"apigateway/internal/bootstrap"
	"context"
)

type StaticPageServiceServer struct {
	staticpagepb.UnimplementedStaticPageServiceServer
	client staticpagepb.StaticPageServiceClient
}

func NewStaticPageServiceServer() *StaticPageServiceServer {
	return &StaticPageServiceServer{
		client: staticpagepb.NewStaticPageServiceClient(bootstrap.MetadataServiceConn),
	}
}

func (s *StaticPageServiceServer) Create(ctx context.Context, req *staticpagepb.CreateRequest) (*staticpagepb.CreateResponse, error) {
	bootstrap.Logger.Info("StaticPageServiceServer: Create called")
	return s.client.Create(ctx, req)
}
