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
	return s.client.Create(ctx, req)
}

func (s *StaticPageServiceServer) Update(ctx context.Context, req *staticpagepb.UpdateRequest) (*staticpagepb.UpdateResponse, error) {
	return s.client.Update(ctx, req)
}

func (s *StaticPageServiceServer) Delete(ctx context.Context, req *staticpagepb.DeleteRequest) (*staticpagepb.DeleteResponse, error) {
	return s.client.Delete(ctx, req)
}

func (s *StaticPageServiceServer) Get(ctx context.Context, req *staticpagepb.GetRequest) (*staticpagepb.GetResponse, error) {
	return s.client.Get(ctx, req)
}

func (s *StaticPageServiceServer) List(ctx context.Context, req *staticpagepb.ListRequest) (*staticpagepb.ListResponse, error) {
	return s.client.List(ctx, req)
}
