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
	resp, err := s.client.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &staticpagepb.CreateResponse{
		Id: resp.Id,
	}, nil
}
